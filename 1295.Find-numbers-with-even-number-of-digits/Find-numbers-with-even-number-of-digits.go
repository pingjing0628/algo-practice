package problem1295

func findNumbers(nums []int) int {
	count, even := 0, 0
	for i := 0; i < len(nums); i++{
		for nums[i] != 0{
			nums[i] = nums[i] / 10
			count++
		}

		if count % 2 == 0 {
			even++
		}

		count = 0
	}

	return even
}
