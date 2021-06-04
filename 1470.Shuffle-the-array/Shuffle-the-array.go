package problem1470

func shuffle(nums []int, n int) []int {

	x := make([]int, 0)
	y := make([]int, 0)
	result := make([]int, 0)

	for i := 0; i < n; i++{
		x = append(x, nums[i])
	}
	for i := n; i < len(nums); i++{
		y = append(y, nums[i])
	}

	for i := 0; i < n; i++{
		result = append(result, x[i])
		result = append(result, y[i])
	}

	return result
}
