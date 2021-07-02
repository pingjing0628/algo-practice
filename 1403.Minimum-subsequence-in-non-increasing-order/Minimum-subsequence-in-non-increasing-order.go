package problem1403

import "sort"

func minSubsequence(nums []int) []int {
	// 由大到小排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	total, subTotal := 0, 0
	result := make([]int, 0)

	for _, n := range nums {
		total += n
	}

	for _, value := range nums {
		subTotal += value
		result = append(result, value)
		// 若 累加值已大於總共減掉累加, 則代表此累加已是最大
		if subTotal > total - subTotal {
			return result
		}
	}

	return result

}
