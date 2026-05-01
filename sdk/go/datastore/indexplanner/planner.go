package indexplanner

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/1backend/1backend/sdk/go/datastore"
)

type branchPlan struct {
	equality map[string]IndexPart
	rng      *IndexPart
}

type shapeSignature struct {
	Branches []shapeBranch `json:"branches,omitempty"`
	Order    []shapeOrder  `json:"order,omitempty"`
	AfterLen int           `json:"afterLen,omitempty"`
}

type shapeBranch struct {
	Filters []shapeFilter `json:"filters,omitempty"`
}

type shapeFilter struct {
	Op        datastore.Op `json:"op"`
	Field     string       `json:"field,omitempty"`
	ValueKind string       `json:"valueKind,omitempty"`
}

type shapeOrder struct {
	Field       string `json:"field,omitempty"`
	Desc        bool   `json:"desc,omitempty"`
	SortingType string `json:"sortingType,omitempty"`
}

func PlanQuery(instance any, q datastore.Query, opts PlanOptions) QueryPlan {
	opts = normalizePlanOptions(opts)
	info := describeModel(instance)
	afterValues := parseAfterValues(q.AfterJson)

	branches, diags := expandFilters(q.Filters, opts.MaxDNFBranches)
	if len(branches) == 0 {
		branches = [][]datastore.Filter{{}}
	}

	shape := shapeSignature{
		AfterLen: len(afterValues),
	}
	for _, orderBy := range q.OrderBys {
		shape.Order = append(shape.Order, shapeOrder{
			Field:       orderBy.Field,
			Desc:        orderBy.Desc,
			SortingType: string(orderBy.SortingType),
		})
	}

	var plans []PlannedIndex
	for _, branch := range branches {
		branchShape, branchPlans, branchDiags := planBranch(info, branch, q.OrderBys, afterValues, opts)
		shape.Branches = append(shape.Branches, branchShape)
		plans = append(plans, branchPlans...)
		diags = append(diags, branchDiags...)
	}

	sort.Slice(shape.Branches, func(i, j int) bool {
		left, _ := json.Marshal(shape.Branches[i])
		right, _ := json.Marshal(shape.Branches[j])
		return string(left) < string(right)
	})

	plans = FinalizePlans(plans)
	return QueryPlan{
		ShapeFingerprint: Fingerprint(shape),
		PlannedIndexes:   plans,
		Diagnostics:      dedupeDiagnostics(diags),
	}
}

func planBranch(info modelInfo, branch []datastore.Filter, orderBys []datastore.OrderBy, afterValues []any, opts PlanOptions) (shapeBranch, []PlannedIndex, []Diagnostic) {
	out := shapeBranch{}
	state := branchPlan{
		equality: map[string]IndexPart{},
	}

	var (
		plans []PlannedIndex
		diags []Diagnostic
	)

	for _, filter := range branch {
		fieldName, valueKind := filterSignature(filter)
		out.Filters = append(out.Filters, shapeFilter{
			Op:        filter.Op,
			Field:     fieldName,
			ValueKind: valueKind,
		})

		if len(filter.Fields) != 1 {
			diags = append(diags, Diagnostic{
				Code:    "unsupported_fields",
				Message: "index planner supports only single-field filters",
			})
			continue
		}
		field, ok := resolveField(info, filter.Fields[0])
		if !ok {
			diags = append(diags, Diagnostic{
				Code:    "unknown_field",
				Message: fmt.Sprintf("field %q cannot be resolved for indexing", filter.Fields[0]),
			})
			continue
		}
		if len(field.Path) > 0 && opts.Dialect == DialectMySQL {
			diags = append(diags, Diagnostic{
				Code:    "unsupported_path_mysql",
				Message: "path-based auto indexes are unsupported for mysql",
			})
			continue
		}

		values := filterValues(filter)
		switch filter.Op {
		case datastore.OpEquals, datastore.OpIsInList:
			if field.Slice {
				if opts.Dialect != DialectPostgres || len(field.Path) > 0 {
					diags = append(diags, Diagnostic{
						Code:    "unsupported_array",
						Message: fmt.Sprintf("field %q cannot be auto-indexed as an array on this backend", filter.Fields[0]),
					})
					continue
				}
				plans = append(plans, PlannedIndex{
					Method: MethodGIN,
					Parts: []IndexPart{{
						Field: field.Field,
					}},
				})
				continue
			}
			if !field.Scalar {
				diags = append(diags, Diagnostic{
					Code:    "unsupported_equality",
					Message: fmt.Sprintf("field %q is not eligible for equality auto indexing", filter.Fields[0]),
				})
				continue
			}
			part := IndexPart{
				Field: field.Field,
				Path:  field.Path,
			}
			state.equality[PathKey(part)] = part
		case datastore.OpLessThan, datastore.OpLessThanOrEqual, datastore.OpGreaterThan, datastore.OpGreaterThanOrEqual:
			if !field.Scalar {
				diags = append(diags, Diagnostic{
					Code:    "unsupported_range",
					Message: fmt.Sprintf("field %q is not eligible for range auto indexing", filter.Fields[0]),
				})
				continue
			}
			part := IndexPart{
				Field: field.Field,
				Path:  field.Path,
				Cast:  comparisonCast(values),
			}
			if state.rng == nil {
				state.rng = &part
			} else if !sameIndexPart(*state.rng, part) {
				diags = append(diags, Diagnostic{
					Code:    "multiple_ranges",
					Message: "index planner supports at most one range field per candidate",
				})
			}
		case datastore.OpIntersects:
			if opts.Dialect != DialectPostgres || !field.Slice || len(field.Path) > 0 {
				diags = append(diags, Diagnostic{
					Code:    "unsupported_intersects",
					Message: fmt.Sprintf("field %q cannot use automatic overlap indexing", filter.Fields[0]),
				})
				continue
			}
			plans = append(plans, PlannedIndex{
				Method: MethodGIN,
				Parts: []IndexPart{{
					Field: field.Field,
				}},
			})
		case datastore.OpContainsSubstring, datastore.OpStartsWith:
			if opts.Dialect != DialectPostgres || !field.Text {
				diags = append(diags, Diagnostic{
					Code:    "unsupported_pattern",
					Message: fmt.Sprintf("field %q cannot use automatic text pattern indexing", filter.Fields[0]),
				})
				continue
			}
			plans = append(plans, PlannedIndex{
				Method:        MethodGIN,
				OperatorClass: "gin_trgm_ops",
				Parts: []IndexPart{{
					Field: field.Field,
					Path:  field.Path,
				}},
			})
		default:
			diags = append(diags, Diagnostic{
				Code:    "unsupported_operator",
				Message: fmt.Sprintf("operator %q is unsupported by the index planner", filter.Op),
			})
		}
	}

	sort.Slice(out.Filters, func(i, j int) bool {
		left, _ := json.Marshal(out.Filters[i])
		right, _ := json.Marshal(out.Filters[j])
		return string(left) < string(right)
	})

	btree, btreeDiags := buildBTreePlan(info, state, orderBys, afterValues)
	diags = append(diags, btreeDiags...)
	if len(btree.Parts) > 0 {
		plans = append(plans, btree)
	}

	return out, plans, diags
}

func buildBTreePlan(info modelInfo, state branchPlan, orderBys []datastore.OrderBy, afterValues []any) (PlannedIndex, []Diagnostic) {
	var diags []Diagnostic
	seen := map[string]struct{}{}
	parts := make([]IndexPart, 0, len(state.equality)+len(orderBys)+1)

	equalities := make([]IndexPart, 0, len(state.equality))
	for _, part := range state.equality {
		equalities = append(equalities, part)
	}
	sort.Slice(equalities, func(i, j int) bool {
		leftRank := schemaRank(info, equalities[i].Field)
		rightRank := schemaRank(info, equalities[j].Field)
		if leftRank != rightRank {
			return leftRank < rightRank
		}
		return PathKey(equalities[i]) < PathKey(equalities[j])
	})

	for _, part := range equalities {
		parts = append(parts, part)
		seen[PathKey(part)+"|"+part.Cast] = struct{}{}
	}

	if state.rng != nil {
		rangeKey := PathKey(*state.rng) + "|" + state.rng.Cast
		if _, ok := seen[rangeKey]; !ok {
			parts = append(parts, *state.rng)
			seen[rangeKey] = struct{}{}
		}
	}

	for i, orderBy := range orderBys {
		if orderBy.SortingType == datastore.SortingTypeRandom {
			diags = append(diags, Diagnostic{
				Code:    "unsupported_random_order",
				Message: "random ordering cannot be auto-indexed",
			})
			continue
		}
		if orderBy.Field == "" {
			diags = append(diags, Diagnostic{
				Code:    "unsupported_anyfield",
				Message: "fieldless ordering cannot be auto-indexed",
			})
			continue
		}

		field, ok := resolveField(info, orderBy.Field)
		if !ok || !field.Scalar {
			diags = append(diags, Diagnostic{
				Code:    "unsupported_order",
				Message: fmt.Sprintf("field %q cannot be auto-indexed for ordering", orderBy.Field),
			})
			continue
		}

		part := IndexPart{
			Field: field.Field,
			Path:  field.Path,
			Desc:  orderBy.Desc,
			Cast:  sortingCast(orderBy, afterValues, i),
		}
		key := PathKey(part) + "|" + part.Cast
		if _, ok := seen[key]; ok {
			continue
		}
		parts = append(parts, part)
		seen[key] = struct{}{}
	}

	if len(parts) == 0 {
		return PlannedIndex{}, diags
	}

	plan := PlannedIndex{
		Method: MethodBTree,
		Parts:  parts,
	}
	plan.Fingerprint = PlanFingerprint(plan)
	return plan, diags
}

func expandFilters(filters []datastore.Filter, max int) ([][]datastore.Filter, []Diagnostic) {
	branches := [][]datastore.Filter{{}}
	var diags []Diagnostic

	for _, filter := range filters {
		options, optionDiags := expandFilter(filter, max)
		diags = append(diags, optionDiags...)
		if len(options) == 0 {
			continue
		}

		nextSize := len(branches) * len(options)
		if nextSize > max {
			diags = append(diags, Diagnostic{
				Code:    "dnf_cap",
				Message: fmt.Sprintf("query exceeded the DNF expansion cap of %d", max),
			})
			return nil, diags
		}

		next := make([][]datastore.Filter, 0, nextSize)
		for _, branch := range branches {
			for _, option := range options {
				merged := append([]datastore.Filter{}, branch...)
				merged = append(merged, option...)
				next = append(next, merged)
			}
		}
		branches = next
	}

	return branches, diags
}

func expandFilter(filter datastore.Filter, max int) ([][]datastore.Filter, []Diagnostic) {
	if filter.Op != datastore.OpOr {
		return [][]datastore.Filter{{filter}}, nil
	}

	if collapsed, ok := collapseEqualityOr(filter); ok {
		return [][]datastore.Filter{{collapsed}}, nil
	}

	var (
		out   [][]datastore.Filter
		diags []Diagnostic
	)
	for _, sub := range filter.SubFilters {
		options, optionDiags := expandFilter(sub, max)
		diags = append(diags, optionDiags...)
		out = append(out, options...)
		if len(out) > max {
			diags = append(diags, Diagnostic{
				Code:    "dnf_cap",
				Message: fmt.Sprintf("query exceeded the DNF expansion cap of %d", max),
			})
			return nil, diags
		}
	}

	return out, diags
}

func collapseEqualityOr(filter datastore.Filter) (datastore.Filter, bool) {
	if filter.Op != datastore.OpOr || len(filter.SubFilters) == 0 {
		return datastore.Filter{}, false
	}

	field := ""
	values := []any{}
	for _, sub := range filter.SubFilters {
		if sub.Op != datastore.OpEquals || len(sub.Fields) != 1 {
			return datastore.Filter{}, false
		}
		if field == "" {
			field = sub.Fields[0]
		}
		if sub.Fields[0] != field {
			return datastore.Filter{}, false
		}
		subValues := filterValues(sub)
		if len(subValues) != 1 {
			return datastore.Filter{}, false
		}
		values = append(values, subValues[0])
	}

	return datastore.IsInList(datastore.Field(field), values...), true
}

func parseAfterValues(raw string) []any {
	if raw == "" {
		return nil
	}

	var values []any
	if err := json.Unmarshal([]byte(raw), &values); err != nil {
		return nil
	}
	return values
}

func filterValues(filter datastore.Filter) []any {
	var values []any
	_ = json.Unmarshal([]byte(filter.ValuesJson), &values)
	return values
}

func comparisonCast(values []any) string {
	if len(values) == 0 {
		return ""
	}
	switch values[0].(type) {
	case float64, float32, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return "numeric"
	default:
		return ""
	}
}

func sortingCast(orderBy datastore.OrderBy, afterValues []any, idx int) string {
	switch orderBy.SortingType {
	case datastore.SortingTypeNumeric:
		return "numeric"
	case datastore.SortingTypeDate:
		return "date"
	case datastore.SortingTypeText:
		return "text"
	}
	if idx < len(afterValues) {
		return comparisonCast([]any{afterValues[idx]})
	}
	return ""
}

func dedupeDiagnostics(diags []Diagnostic) []Diagnostic {
	out := make([]Diagnostic, 0, len(diags))
	seen := map[string]struct{}{}
	for _, diag := range diags {
		key := diag.Code + "|" + diag.Message
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, diag)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Code != out[j].Code {
			return out[i].Code < out[j].Code
		}
		return out[i].Message < out[j].Message
	})
	return out
}

func filterSignature(filter datastore.Filter) (string, string) {
	field := strings.Join(filter.Fields, ",")
	if len(filter.Fields) == 0 {
		field = "*"
	}
	values := filterValues(filter)
	if len(values) == 0 {
		return field, ""
	}

	switch v := values[0].(type) {
	case string:
		return field, "string"
	case bool:
		return field, "bool"
	case float64, float32, int, int8, int16, int32, int64:
		return field, "number"
	case []any:
		_ = v
		return field, "array"
	default:
		return field, fmt.Sprintf("%T", v)
	}
}
