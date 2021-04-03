package permutations

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "测试正常的全排列",
			nums: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
