package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	ts := []struct {
		Input string
		Want  bool
	}{
		{
			Input: "()",
			Want:  true,
		},
		{
			Input: "()[]{}",
			Want:  true,
		},
		{
			Input: "(]",
			Want:  false,
		},
		{
			Input: "{[]}",
			Want:  true,
		},
		{
			Input: "{[}]",
			Want:  false,
		},
		{
			Input: "}",
			Want:  false,
		},
	}

	for _, tc := range ts {
		got := isValid(tc.Input)
		if got != tc.Want {
			t.Errorf("failed. input: %v, got: %v, want: %v",
				tc.Input, got, tc.Want,
			)
		}
	}
}

func TestStack(t *testing.T) {
	s := NewStack(4)

	if s.String() != "" {
		t.Errorf("should be empty stack: %s", s)
	}

	s1 := "hello"
	for _, c := range s1 {
		s.push(byte(c))
	}
	if s.String() != s1 {
		t.Errorf("should be %s, but got: %s", s1, s)
	}

	s.pop()
	if s.String() != s1[:len(s1)-1] {
		t.Errorf("should be %s, but got: %s", s1[:len(s1)-1], s)
	}
}
