package indexplanner

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/stretchr/testify/require"
)

type plannerFriend struct {
	Age int `json:"age"`
}

type plannerObject struct {
	Id        string        `json:"id"`
	Name      string        `json:"name"`
	Value     int           `json:"value"`
	CreatedAt time.Time     `json:"createdAt"`
	Tags      []string      `json:"tags"`
	Friend    plannerFriend `json:"friend"`
}

func (p plannerObject) GetId() string {
	return p.Id
}

func TestPlanQueryBTree(t *testing.T) {
	after, err := json.Marshal([]any{"2025-01-01T00:00:00Z"})
	require.NoError(t, err)

	plan := PlanQuery(plannerObject{}, datastore.Query{
		Filters: []datastore.Filter{
			datastore.Equals(datastore.Field("name"), "alice"),
			datastore.GreaterThan(datastore.Field("value"), 10),
		},
		OrderBys: []datastore.OrderBy{
			datastore.OrderByField("createdAt", true),
		},
		AfterJson: string(after),
	}, PlanOptions{Dialect: DialectPostgres})

	require.Empty(t, plan.Diagnostics)
	require.Len(t, plan.PlannedIndexes, 1)

	idx := plan.PlannedIndexes[0]
	require.Equal(t, MethodBTree, idx.Method)
	require.Equal(t, []IndexPart{
		{Field: "name"},
		{Field: "value", Cast: "numeric"},
		{Field: "createdAt", Desc: true},
	}, idx.Parts)
}

func TestPlanQueryPathArrayAndTrigram(t *testing.T) {
	t.Run("path", func(t *testing.T) {
		plan := PlanQuery(plannerObject{}, datastore.Query{
			Filters: []datastore.Filter{
				datastore.LessThan(datastore.Field("friend.age"), 30),
			},
		}, PlanOptions{Dialect: DialectPostgres})

		require.Len(t, plan.PlannedIndexes, 1)
		require.Equal(t, PlannedIndex{
			Method:      MethodBTree,
			Parts:       []IndexPart{{Field: "friend", Path: []string{"age"}, Cast: "numeric"}},
			Fingerprint: plan.PlannedIndexes[0].Fingerprint,
		}, plan.PlannedIndexes[0])
	})

	t.Run("array", func(t *testing.T) {
		plan := PlanQuery(plannerObject{}, datastore.Query{
			Filters: []datastore.Filter{
				datastore.Intersects(datastore.Field("tags"), []any{"tag-a"}),
			},
		}, PlanOptions{Dialect: DialectPostgres})

		require.Len(t, plan.PlannedIndexes, 1)
		require.Equal(t, MethodGIN, plan.PlannedIndexes[0].Method)
		require.Equal(t, []IndexPart{{Field: "tags"}}, plan.PlannedIndexes[0].Parts)
	})

	t.Run("trigram", func(t *testing.T) {
		plan := PlanQuery(plannerObject{}, datastore.Query{
			Filters: []datastore.Filter{
				datastore.ContainsSubstring(datastore.Field("name"), "ali"),
			},
		}, PlanOptions{Dialect: DialectPostgres})

		require.Len(t, plan.PlannedIndexes, 1)
		require.Equal(t, MethodGIN, plan.PlannedIndexes[0].Method)
		require.Equal(t, "gin_trgm_ops", plan.PlannedIndexes[0].OperatorClass)
		require.Equal(t, []IndexPart{{Field: "name"}}, plan.PlannedIndexes[0].Parts)
	})
}

func TestPlanQueryOrCollapseAndCap(t *testing.T) {
	t.Run("collapse equality or", func(t *testing.T) {
		plan := PlanQuery(plannerObject{}, datastore.Query{
			Filters: []datastore.Filter{
				datastore.Or(
					datastore.Equals(datastore.Field("name"), "alice"),
					datastore.Equals(datastore.Field("name"), "bob"),
				),
			},
		}, PlanOptions{Dialect: DialectPostgres})

		require.Empty(t, plan.Diagnostics)
		require.Len(t, plan.PlannedIndexes, 1)
		require.Equal(t, []IndexPart{{Field: "name"}}, plan.PlannedIndexes[0].Parts)
	})

	t.Run("dnf cap", func(t *testing.T) {
		plan := PlanQuery(plannerObject{}, datastore.Query{
			Filters: []datastore.Filter{
				datastore.Or(datastore.Equals(datastore.Field("name"), "alice"), datastore.Equals(datastore.Field("value"), 1)),
				datastore.Or(datastore.Equals(datastore.Field("friend.age"), 3), datastore.ContainsSubstring(datastore.Field("name"), "al")),
				datastore.Or(datastore.Equals(datastore.Field("createdAt"), "2025-01-01T00:00:00Z"), datastore.Equals(datastore.Field("tags"), "tag-a")),
			},
		}, PlanOptions{
			Dialect:        DialectPostgres,
			MaxDNFBranches: 4,
		})

		require.Empty(t, plan.PlannedIndexes)
		require.NotEmpty(t, plan.Diagnostics)
		require.Equal(t, "dnf_cap", plan.Diagnostics[0].Code)
	})
}

func TestTrackerPromotionCoverageAndCap(t *testing.T) {
	plan := PlanQuery(plannerObject{}, datastore.Query{
		Filters: []datastore.Filter{
			datastore.Equals(datastore.Field("name"), "alice"),
		},
	}, PlanOptions{Dialect: DialectPostgres})
	require.Len(t, plan.PlannedIndexes, 1)

	tracker := NewTracker(TrackerOptions{
		Backend:            DialectPostgres,
		Supported:          true,
		PromotionThreshold: 3,
		MaxAutoIndexes:     1,
	})

	manual := datastore.Index{Fields: []string{"name", "createdAt"}, Ascending: true}
	tracker.RegisterManual(ManualIndexPlan(manual), "manual_name_createdAt_idx")

	require.Empty(t, tracker.Observe(plan).Candidates)
	require.Empty(t, tracker.Observe(plan).Candidates)
	require.Empty(t, tracker.Observe(plan).Candidates)

	stats := tracker.Stats()
	require.Len(t, stats.Shapes, 1)
	require.Equal(t, 3, stats.Shapes[0].Hits)
	require.Equal(t, "covered by existing index", stats.Shapes[0].Reason)

	secondPlan := PlanQuery(plannerObject{}, datastore.Query{
		Filters: []datastore.Filter{
			datastore.Equals(datastore.Field("value"), 42),
		},
	}, PlanOptions{Dialect: DialectPostgres})

	require.Empty(t, tracker.Observe(secondPlan).Candidates)
	require.Empty(t, tracker.Observe(secondPlan).Candidates)

	result := tracker.Observe(secondPlan)
	require.Len(t, result.Candidates, 1)

	candidate := result.Candidates[0]
	require.True(t, tracker.MarkPending(candidate, "table_autoidx_"+candidate.Fingerprint))
	require.False(t, tracker.MarkPending(candidate, "table_autoidx_"+candidate.Fingerprint))
	tracker.MarkCreated(candidate, "table_autoidx_"+candidate.Fingerprint)

	thirdPlan := PlanQuery(plannerObject{}, datastore.Query{
		Filters: []datastore.Filter{
			datastore.Equals(datastore.Field("friend.age"), 9),
		},
	}, PlanOptions{Dialect: DialectPostgres})

	require.Empty(t, tracker.Observe(thirdPlan).Candidates)
	require.Empty(t, tracker.Observe(thirdPlan).Candidates)
	require.Empty(t, tracker.Observe(thirdPlan).Candidates)

	stats = tracker.Stats()
	auto := []datastore.AutoIndexEntry{}
	for _, entry := range stats.Indexes {
		if entry.Kind == datastore.AutoIndexKindAuto {
			auto = append(auto, entry)
		}
	}
	require.Len(t, auto, 1)
	require.Equal(t, datastore.AutoIndexStatusCreated, auto[0].Status)
}
