package merge_orted_array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args

		want []int
	}{
		{
			name: "示例1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "示例2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			want: []int{1},
		},
		{
			name: "示例3",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1},
		},
		{
			name: "示例4",
			args: args{
				nums1: []int{1, 0},
				m:     1,
				nums2: []int{2},
				n:     1,
			},
			want: []int{1, 2},
		},
		{
			name: "示例5",
			args: args{
				nums1: []int{2, 0},
				m:     1,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums1 := tt.args.nums1
			merge(nums1, tt.args.m, tt.args.nums2, tt.args.n)

			sort.Ints(nums1)
			sort.Ints(tt.want)

			assert.Equal(t, tt.want, nums1)
		})
	}
}
