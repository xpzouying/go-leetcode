package letter_case_permutation

import "testing"

func TestLetterCasePermutation(t *testing.T) {

	ts := []struct {
		Input string
		Want  []string
	}{
		{
			Input: "a1b2",
			Want:  []string{"a1b2", "a1B2", "A1b2", "A1B2"},
		},
	}

	for _, tc := range ts {
		got := letterCasePermutation(tc.Input)
		if !sameList(tc.Want, got) {
			t.Errorf("want=%+v\ngot=%+v", tc.Want, got)
		}
	}
}

func sameList(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	ma := make(map[string]bool, len(a))
	for _, v := range a {
		ma[v] = true
	}

	for _, v := range b {
		if _, ok := ma[v]; !ok {
			return false
		}
	}

	return true
}
