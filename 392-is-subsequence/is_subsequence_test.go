package is_usbsequence

import "testing"

func TestIsSubsequence(t *testing.T) {

	ts := []struct {
		slist string
		tlist string
		want  bool
	}{
		{
			slist: "abc",
			tlist: "ahbgdc",
			want:  true,
		},
		{
			slist: "axc",
			tlist: "ahbgdc",
			want:  false,
		},
		{
			slist: "a",
			tlist: "a",
			want:  true,
		},
	}

	for _, tc := range ts {
		got := isSubsequence(tc.slist, tc.tlist)
		if tc.want != got {
			t.Errorf("sub_list: %s, target_list: %s, want: %v, got: %v",
				tc.slist, tc.tlist, tc.want, got)
		}
	}

}
