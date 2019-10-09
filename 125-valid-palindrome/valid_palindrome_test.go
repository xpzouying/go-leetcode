package valid_palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {

	ts := []struct {
		input string
		want  bool
	}{
		{
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},
		{
			input: "race a car",
			want:  false,
		},
		{
			input: "race e c a r",
			want:  true,
		},
		{
			input: "",
			want:  true,
		},
		{
			input: "0P",
			want:  false,
		},
	}

	for _, tc := range ts {
		got := isPalindrome(tc.input)
		if tc.want != got {
			t.Errorf("input:%s, is palindrome:%v, got:%v", tc.input, tc.want, got)
		}
	}

}
