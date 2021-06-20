package problem1464

import "sort"

func maxProduct(nums []int) int {
	// Solution 1
	sort.Ints(nums)
	return (nums[len(nums) - 1] - 1) * (nums[len(nums) - 1 - 1] - 1)


	// Solution 2
	// a, b := 0, 0
	// for _, v := range nums {
	//     if v > a {
	//         a, b = v, a
	//     } else if v > b {
	//         b = v
	//     }
	// }
	// return (a - 1) * (b - 1)
}
