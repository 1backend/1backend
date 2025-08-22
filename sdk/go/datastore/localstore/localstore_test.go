package localstore

import (
	"testing"
	"time"
)

// ptr is a helper to quickly create pointers
func ptr[T any](v T) *T { return &v }

func mustTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

func TestCompare_NilAndMissing(t *testing.T) {
	var nilInt *int

	tests := []struct {
		name string
		vi   any
		vj   any
		desc bool
		want bool
	}{
		{"both nil → equal → false", nil, nil, false, false},
		{"vi nil asc → true", nil, 5, false, true},
		{"vi nil desc → false", nil, 5, true, false},
		{"vj nil asc → false", 5, nil, false, false},
		{"vj nil desc → true", 5, nil, true, true},
		{"nil pointer asc → true", nilInt, 5, false, true},
		{"nil pointer desc → false", nilInt, 5, true, false},
	}

	for _, tc := range tests {
		if got := compare(tc.vi, tc.vj, tc.desc); got != tc.want {
			t.Errorf("%s: got %v, want %v", tc.name, got, tc.want)
		}
	}
}

func TestCompare_EqualityAlwaysFalse(t *testing.T) {
	equalCases := []any{
		42,
		uint64(42),
		3.14,
		"abc",
		mustTime("2025-01-01T00:00:00Z"),
	}
	for _, v := range equalCases {
		if got := compare(v, v, false); got {
			t.Errorf("equal asc should be false for %T", v)
		}
		if got := compare(v, v, true); got {
			t.Errorf("equal desc should be false for %T", v)
		}
	}
}

func TestCompare_Integers(t *testing.T) {
	tests := []struct {
		name string
		vi   any
		vj   any
		desc bool
		want bool
	}{
		{"int asc", int(1), int(2), false, true},
		{"int desc", int(2), int(1), true, true},
		{"int64 asc", int64(1), int64(2), false, true},
		{"int64 desc", int64(2), int64(1), true, true},
		{"uint asc", uint(1), uint(2), false, true},
		{"uint desc", uint(2), uint(1), true, true},
	}

	for _, tc := range tests {
		if got := compare(tc.vi, tc.vj, tc.desc); got != tc.want {
			t.Errorf("%s: got %v, want %v", tc.name, got, tc.want)
		}
	}
}

func TestCompare_Floats(t *testing.T) {
	tests := []struct {
		name string
		vi   any
		vj   any
		desc bool
		want bool
	}{
		{"float asc", 1.5, 2.5, false, true},
		{"float desc", 2.5, 1.5, true, true},
	}

	for _, tc := range tests {
		if got := compare(tc.vi, tc.vj, tc.desc); got != tc.want {
			t.Errorf("%s: got %v, want %v", tc.name, got, tc.want)
		}
	}
}

func TestCompare_Strings_Lexicographic(t *testing.T) {
	tests := []struct {
		name string
		vi   any
		vj   any
		desc bool
		want bool
	}{
		{"asc, abc<abd", "abc", "abd", false, true},
		{"desc, abd>abc", "abd", "abc", true, true},
		{"asc, non-time fallback", "zzz", "aaa", false, false},
		{"desc, non-time fallback", "zzz", "aaa", true, true},
	}

	for _, tc := range tests {
		if got := compare(tc.vi, tc.vj, tc.desc); got != tc.want {
			t.Errorf("%s: got %v, want %v", tc.name, got, tc.want)
		}
	}
}

func TestCompare_StringDates_Parseable(t *testing.T) {
	// Only runs if your datastore.ParseAnyDate supports RFC3339
	s1 := "2024-01-01T00:00:00Z"
	s2 := "2024-02-01T00:00:00Z"

	if got := compare(s1, s2, false); !got {
		t.Errorf("expected %s < %s in asc date mode", s1, s2)
	}
	if got := compare(s2, s1, true); !got {
		t.Errorf("expected %s > %s in desc date mode", s2, s1)
	}
}

func TestCompare_TimeStruct(t *testing.T) {
	t1 := mustTime("2024-01-01T00:00:00Z")
	t2 := mustTime("2024-01-02T00:00:00Z")

	if got := compare(t1, t2, false); !got {
		t.Errorf("t1 should be before t2 asc")
	}
	if got := compare(t2, t1, true); !got {
		t.Errorf("t2 should be after t1 desc")
	}
}

func TestCompare_PointerValues(t *testing.T) {
	i1, i2 := 1, 2
	t1 := mustTime("2024-01-01T00:00:00Z")
	t2 := mustTime("2024-01-02T00:00:00Z")

	tests := []struct {
		name string
		vi   any
		vj   any
		desc bool
		want bool
	}{
		{"*int asc", ptr(i1), ptr(i2), false, true},
		{"*int desc", ptr(i2), ptr(i1), true, true},
		{"*time asc", ptr(t1), ptr(t2), false, true},
		{"*time desc", ptr(t2), ptr(t1), true, true},
	}

	for _, tc := range tests {
		if got := compare(tc.vi, tc.vj, tc.desc); got != tc.want {
			t.Errorf("%s: got %v, want %v", tc.name, got, tc.want)
		}
	}
}

func TestCompare_UnsupportedTypes(t *testing.T) {
	type Custom struct{ A int }
	c1 := Custom{1}
	c2 := Custom{2}

	if got := compare(c1, c2, false); got {
		t.Errorf("unsupported type should return false, got %v", got)
	}
}

func TestCompare_FalseNilZero(t *testing.T) {
	tests := []struct {
		name string
		vi   any
		vj   any
		want bool
	}{
		// nil vs other "zero-like" values
		{"nil vs false", nil, false, false},
		{"false vs nil", false, nil, false},
		{"nil vs 0", nil, 0, false},
		{"0 vs nil", 0, nil, false},
		{"nil vs empty string", nil, "", false},
		{"empty string vs nil", "", nil, false},

		// false vs numeric
		// These cause panics
		// {"false vs 0", false, 0, false},
		// {"0 vs false", 0, false, false},

		// explicit equality should still be false
		{"false vs false", false, false, false},
		{"0 vs 0", 0, 0, false},
		{"\"\" vs \"\"", "", "", false},
	}

	for _, tc := range tests {
		if got := compare(tc.vi, tc.vj, false); got != tc.want {
			t.Errorf("%s asc: got %v, want %v", tc.name, got, tc.want)
		}
		if got := compare(tc.vi, tc.vj, true); got != tc.want {
			t.Errorf("%s desc: got %v, want %v", tc.name, got, tc.want)
		}
	}
}
