package range_sum_query_immutable

//// type NumArray struct {
////     Sum [][]int  // Sum[i][j]
//// }
////
////
//// func Constructor(nums []int) NumArray {
//// /**
////
//// num(i, j) = num(i, j-1) + A[j]
////
////
//// num(0, 0) = -2
//// num(0, 1) = num(0, j-1) + A(j) = num(0, 0) + 0 = -2 + 0 = -2
////
////
////
//// */
////
////     if len(nums) == 0 {
////         return NumArray{
////             Sum: nil,
////         }
////     }
////
////     // use 2-d to store the record
////     // Sum[i][j] is sum of begin i to end j
////
////     Sum := make([][]int, len(nums))
////     /*
////     for i := 0; i < len(nums); i++ {
////         Sum[i] = make([]int, len(nums))
////     }
////     Sum[0][0] = nums[0]
////     */
////
////     for i := 0; i < len(nums); i++ {
////         Sum[i] = make([]int, len(nums))
////
////         for j := i; j < len(nums); j++ {
////             if i == j {
////                 Sum[i][j] = nums[i]
////                 continue
////             }
////
////             Sum[i][j] = Sum[i][j-1] + nums[j]
////         }
////     }
////
////     return NumArray{
////         Sum: Sum,
////     }
//// }
////
////
//// func (this *NumArray) SumRange(i int, j int) int {
////     if i > j {
////         return 0
////     }
////
////     if j > len(this.Sum)-1 {
////         return 0
////     }
////
////     return this.Sum[i][j]
//// }
////
////
//// /**
////  * Your NumArray object will be instantiated and called as such:
////  * obj := Constructor(nums);
////  * param_1 := obj.SumRange(i,j);
////  */

// --- dynamic programming ---
//
/**

Notebook:
1, 1, 2, 5, 10

(0, 0): 1
(0, 1):
sum[0] = 1
sum[1] = 2
sum[2] = 4
sum[3] = 4+5 = 9
sum[4] = 9+10 = 19

(0, 1) = 2
(1, 1) = 1
(2, 4) = 9-2 = 7
(3, 4) = sum(4) - sum(2) = 19-4 = 15
*/

type NumArray struct {
	Sum []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{
			Sum: nil,
		}
	}

	Sum := make([]int, len(nums))
	Sum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		Sum[i] = Sum[i-1] + nums[i]
	}

	return NumArray{
		Sum: Sum,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	Sum := this.Sum

	if j > len(Sum)-1 {
		return 0
	}

	if i == 0 {
		return Sum[j]
	}

	return Sum[j] - Sum[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
