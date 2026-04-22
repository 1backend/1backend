package indexplanner

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
)

const (
	DialectGeneric  = "generic"
	DialectPostgres = "postgres"
	DialectMySQL    = "mysql"

	MethodBTree = "btree"
	MethodGIN   = "gin"

	DefaultPromotionThreshold = 3
	DefaultMaxAutoIndexes     = 8
	DefaultDNFExpansionCap    = 16
)

type PlanOptions struct {
	Dialect        string
	MaxDNFBranches int
}

type QueryPlan struct {
	ShapeFingerprint string
	PlannedIndexes   []PlannedIndex
	Diagnostics      []Diagnostic
}

type PlannedIndex struct {
	Method        string
	Parts         []IndexPart
	OperatorClass string
	Unique        bool
	Fingerprint   string
}

type IndexPart struct {
	Field string
	Path  []string
	Desc  bool
	Cast  string
}

type Diagnostic struct {
	Code    string
	Message string
}

type TrackerOptions struct {
	Backend            string
	Supported          bool
	PromotionThreshold int
	MaxAutoIndexes     int
}

type ObserveResult struct {
	Candidates []PlannedIndex
}

type Tracker struct {
	mu sync.Mutex

	opts TrackerOptions

	shapes         map[string]*datastore.AutoIndexShapeStats
	entries        map[string]*trackedEntry
	manualPlans    []PlannedIndex
	autoIndexCount int
}

type trackedEntry struct {
	Plan  PlannedIndex
	Entry datastore.AutoIndexEntry
}

func normalizePlanOptions(opts PlanOptions) PlanOptions {
	if opts.Dialect == "" {
		opts.Dialect = DialectGeneric
	}
	if opts.MaxDNFBranches <= 0 {
		opts.MaxDNFBranches = DefaultDNFExpansionCap
	}
	return opts
}

func normalizeTrackerOptions(opts TrackerOptions) TrackerOptions {
	if opts.PromotionThreshold <= 0 {
		opts.PromotionThreshold = DefaultPromotionThreshold
	}
	if opts.MaxAutoIndexes <= 0 {
		opts.MaxAutoIndexes = DefaultMaxAutoIndexes
	}
	return opts
}

func NewTracker(opts TrackerOptions) *Tracker {
	opts = normalizeTrackerOptions(opts)
	return &Tracker{
		opts:    opts,
		shapes:  map[string]*datastore.AutoIndexShapeStats{},
		entries: map[string]*trackedEntry{},
	}
}

func Fingerprint(v any) string {
	bs, _ := json.Marshal(v)
	sum := sha1.Sum(bs)
	return hex.EncodeToString(sum[:])[:32]
}

func FinalizePlans(plans []PlannedIndex) []PlannedIndex {
	out := make([]PlannedIndex, 0, len(plans))
	seen := map[string]struct{}{}

	for _, plan := range plans {
		plan.Fingerprint = PlanFingerprint(plan)
		if _, ok := seen[plan.Fingerprint]; ok {
			continue
		}
		seen[plan.Fingerprint] = struct{}{}
		out = append(out, plan)
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Fingerprint < out[j].Fingerprint
	})

	return out
}

func PlanFingerprint(plan PlannedIndex) string {
	type part struct {
		Field string   `json:"field"`
		Path  []string `json:"path,omitempty"`
		Desc  bool     `json:"desc,omitempty"`
		Cast  string   `json:"cast,omitempty"`
	}

	parts := make([]part, 0, len(plan.Parts))
	for _, p := range plan.Parts {
		parts = append(parts, part{
			Field: p.Field,
			Path:  append([]string(nil), p.Path...),
			Desc:  p.Desc,
			Cast:  p.Cast,
		})
	}

	return Fingerprint(struct {
		Method        string `json:"method"`
		OperatorClass string `json:"operatorClass,omitempty"`
		Unique        bool   `json:"unique,omitempty"`
		Parts         []part `json:"parts"`
	}{
		Method:        plan.Method,
		OperatorClass: plan.OperatorClass,
		Unique:        plan.Unique,
		Parts:         parts,
	})
}

func PathKey(part IndexPart) string {
	if len(part.Path) == 0 {
		return part.Field
	}
	return part.Field + "." + strings.Join(part.Path, ".")
}

func Covers(existing, candidate PlannedIndex) bool {
	if existing.Method != candidate.Method || existing.OperatorClass != candidate.OperatorClass || existing.Unique != candidate.Unique {
		return false
	}

	switch existing.Method {
	case MethodBTree:
		if len(existing.Parts) < len(candidate.Parts) {
			return false
		}
		for i := range candidate.Parts {
			if !sameIndexPart(existing.Parts[i], candidate.Parts[i]) {
				return false
			}
		}
		return true
	default:
		return existing.Fingerprint == candidate.Fingerprint
	}
}

func ManualIndexPlan(index datastore.Index) PlannedIndex {
	parts := make([]IndexPart, 0, len(index.Fields))
	for _, field := range index.Fields {
		parts = append(parts, IndexPart{
			Field: field,
			Desc:  !index.Ascending,
		})
	}

	plan := PlannedIndex{
		Method: MethodBTree,
		Parts:  parts,
		Unique: index.Unique,
	}
	plan.Fingerprint = PlanFingerprint(plan)
	return plan
}

func (t *Tracker) RegisterManual(plan PlannedIndex, name string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	plan.Fingerprint = ensurePlanFingerprint(plan)
	t.manualPlans = append(t.manualPlans, plan)
	t.entries[plan.Fingerprint] = &trackedEntry{
		Plan: plan,
		Entry: datastore.AutoIndexEntry{
			Fingerprint: plan.Fingerprint,
			Kind:        datastore.AutoIndexKindManual,
			Method:      plan.Method,
			Status:      datastore.AutoIndexStatusManual,
			Name:        name,
		},
	}
}

func (t *Tracker) RegisterDiscovered(fingerprint, method, name string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, ok := t.entries[fingerprint]; ok {
		return
	}

	t.entries[fingerprint] = &trackedEntry{
		Plan: PlannedIndex{
			Method:      method,
			Fingerprint: fingerprint,
		},
		Entry: datastore.AutoIndexEntry{
			Fingerprint: fingerprint,
			Kind:        datastore.AutoIndexKindAuto,
			Method:      method,
			Status:      datastore.AutoIndexStatusDiscovered,
			Name:        name,
		},
	}
	t.autoIndexCount++
}

func (t *Tracker) Observe(plan QueryPlan) ObserveResult {
	t.mu.Lock()
	defer t.mu.Unlock()

	now := time.Now().UTC()

	state := t.shapes[plan.ShapeFingerprint]
	if state == nil {
		state = &datastore.AutoIndexShapeStats{
			Fingerprint: plan.ShapeFingerprint,
		}
		t.shapes[plan.ShapeFingerprint] = state
	}

	state.Hits++
	state.LastSeen = now
	state.Eligible = len(plan.PlannedIndexes) > 0 && t.opts.Supported
	state.Reason = trackerReason(plan.Diagnostics)

	if !t.opts.Supported {
		if state.Reason == "" {
			state.Reason = "auto indexing unsupported by backend"
		}
		return ObserveResult{}
	}

	if state.Hits < t.opts.PromotionThreshold {
		if state.Reason == "" && len(plan.PlannedIndexes) > 0 {
			state.Reason = fmt.Sprintf("waiting for %d hits", t.opts.PromotionThreshold)
		}
		return ObserveResult{}
	}

	res := ObserveResult{}
	for _, candidate := range plan.PlannedIndexes {
		candidate.Fingerprint = ensurePlanFingerprint(candidate)
		if t.isCovered(candidate) {
			if state.Reason == "" {
				state.Reason = "covered by existing index"
			}
			continue
		}
		if t.autoIndexCount >= t.opts.MaxAutoIndexes {
			state.Reason = "auto index cap reached"
			continue
		}
		res.Candidates = append(res.Candidates, candidate)
	}

	if len(res.Candidates) > 0 {
		state.Reason = ""
	}

	return res
}

func (t *Tracker) MarkPending(plan PlannedIndex, name string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	plan.Fingerprint = ensurePlanFingerprint(plan)
	if t.isCovered(plan) || t.autoIndexCount >= t.opts.MaxAutoIndexes {
		return false
	}

	t.entries[plan.Fingerprint] = &trackedEntry{
		Plan: plan,
		Entry: datastore.AutoIndexEntry{
			Fingerprint: plan.Fingerprint,
			Kind:        datastore.AutoIndexKindAuto,
			Method:      plan.Method,
			Status:      datastore.AutoIndexStatusPending,
			Name:        name,
			LastSeen:    time.Now().UTC(),
		},
	}
	t.autoIndexCount++
	return true
}

func (t *Tracker) MarkCreated(plan PlannedIndex, name string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	plan.Fingerprint = ensurePlanFingerprint(plan)
	entry, ok := t.entries[plan.Fingerprint]
	if !ok {
		t.entries[plan.Fingerprint] = &trackedEntry{
			Plan: plan,
			Entry: datastore.AutoIndexEntry{
				Fingerprint: plan.Fingerprint,
				Kind:        datastore.AutoIndexKindAuto,
				Method:      plan.Method,
				Status:      datastore.AutoIndexStatusCreated,
				Name:        name,
				LastSeen:    time.Now().UTC(),
			},
		}
		t.autoIndexCount++
		return
	}

	entry.Plan = plan
	entry.Entry.Method = plan.Method
	entry.Entry.Status = datastore.AutoIndexStatusCreated
	entry.Entry.Name = name
	entry.Entry.LastSeen = time.Now().UTC()
	entry.Entry.Error = ""
}

func (t *Tracker) MarkFailed(plan PlannedIndex, name string, err error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	plan.Fingerprint = ensurePlanFingerprint(plan)
	entry, ok := t.entries[plan.Fingerprint]
	if !ok {
		return
	}

	if entry.Entry.Kind == datastore.AutoIndexKindAuto && entry.Entry.Status == datastore.AutoIndexStatusPending && t.autoIndexCount > 0 {
		t.autoIndexCount--
	}

	entry.Plan = plan
	entry.Entry.Method = plan.Method
	entry.Entry.Status = datastore.AutoIndexStatusFailed
	entry.Entry.Name = name
	entry.Entry.LastSeen = time.Now().UTC()
	if err != nil {
		entry.Entry.Error = err.Error()
	}
}

func (t *Tracker) Stats() datastore.AutoIndexStats {
	t.mu.Lock()
	defer t.mu.Unlock()

	stats := datastore.AutoIndexStats{
		Supported: t.opts.Supported,
		Backend:   t.opts.Backend,
	}

	for _, shape := range t.shapes {
		stats.Shapes = append(stats.Shapes, *shape)
	}
	for _, entry := range t.entries {
		stats.Indexes = append(stats.Indexes, entry.Entry)
	}

	sort.Slice(stats.Shapes, func(i, j int) bool {
		return stats.Shapes[i].Fingerprint < stats.Shapes[j].Fingerprint
	})
	sort.Slice(stats.Indexes, func(i, j int) bool {
		if stats.Indexes[i].Kind != stats.Indexes[j].Kind {
			return stats.Indexes[i].Kind < stats.Indexes[j].Kind
		}
		return stats.Indexes[i].Fingerprint < stats.Indexes[j].Fingerprint
	})

	return stats
}

func (t *Tracker) isCovered(candidate PlannedIndex) bool {
	candidate.Fingerprint = ensurePlanFingerprint(candidate)

	for _, manual := range t.manualPlans {
		if Covers(manual, candidate) {
			return true
		}
	}
	for _, entry := range t.entries {
		if entry.Entry.Kind != datastore.AutoIndexKindAuto {
			continue
		}
		if entry.Entry.Status == datastore.AutoIndexStatusFailed {
			continue
		}
		if entry.Plan.Fingerprint == "" {
			if entry.Entry.Fingerprint == candidate.Fingerprint {
				return true
			}
			continue
		}
		if Covers(entry.Plan, candidate) {
			return true
		}
	}
	return false
}

func trackerReason(diags []Diagnostic) string {
	if len(diags) == 0 {
		return ""
	}
	return diags[0].Message
}

func ensurePlanFingerprint(plan PlannedIndex) string {
	if plan.Fingerprint != "" {
		return plan.Fingerprint
	}
	return PlanFingerprint(plan)
}

func sameIndexPart(left, right IndexPart) bool {
	if left.Field != right.Field || left.Desc != right.Desc || left.Cast != right.Cast {
		return false
	}
	if len(left.Path) != len(right.Path) {
		return false
	}
	for i := range left.Path {
		if left.Path[i] != right.Path[i] {
			return false
		}
	}
	return true
}
