package find_disappeared_numbers

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			"示例1",
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDisappearedNumbers(tt.nums)

			sort.Ints(got)

			want := tt.want
			sort.Ints(want)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, want)
			}
		})
	}
}
