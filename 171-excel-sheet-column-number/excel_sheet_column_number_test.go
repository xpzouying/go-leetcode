package excel_sheet_column_number

import "testing"

func TestTitleToNumber(t *testing.T) {

	ts := []struct {
		Input string
		Want  int
	}{
		{Input: "A", Want: 1},
		{Input: "AB", Want: 28},
		{Input: "ZY", Want: 701},
	}

	for _, tc := range ts {
		got := titleToNumber(tc.Input)
		if tc.Want != got {
			t.Errorf("Input:%v, Want:%v, Got:%v", tc.Input, tc.Want, got)
		}
	}

}
