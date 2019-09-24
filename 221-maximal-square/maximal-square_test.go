package maximal_square

import "testing"

func TestMaxMatrix(t *testing.T) {

	ts := []struct {
		matrix [][]byte
		want   int
	}{
		{
			matrix: [][]byte{
				[]byte("10100"),
				[]byte("10111"),
				[]byte("11111"),
				[]byte("10010"),
			},
			want: 4,
		},
		{
			matrix: [][]byte{
				[]byte("111"),
				[]byte("111"),
				[]byte("111"),
			},
			want: 9,
		},
		{
			matrix: [][]byte{
				[]byte{'1'},
			},
			want: 1,
		},
	}

	for _, tc := range ts {
		got := maximalSquare(tc.matrix)
		if tc.want != got {
			t.Errorf("matrix:%+v, want:%d, got:%d",
				tc.matrix, tc.want, got)
		}
	}
}
