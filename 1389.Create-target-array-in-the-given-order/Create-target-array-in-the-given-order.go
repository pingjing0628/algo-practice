package problem1389

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, len(nums))
	for i, value := range index {
		copy(target[value + 1:], target[value:])
		target[value] = nums[i]

	}

	return target
}
