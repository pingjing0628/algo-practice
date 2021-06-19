package problem1827

func minOperations(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++{
		if nums[i] <= nums[i - 1] {
			count += nums[i - 1] - nums[i] + 1
			nums[i] += nums[i - 1] - nums[i] + 1
		}
	}

	return count
}
