package problem1304

func sumZero(n int) []int {
	result := make([]int, 0)
	for i := 1; i < n / 2 + 1; i++{
		result = append(result, []int{i, -i}...)
		// result = append(result, i, -i)
	}

	if n % 2 != 0 {
		result = append(result, 0)
	}

	return result
}
