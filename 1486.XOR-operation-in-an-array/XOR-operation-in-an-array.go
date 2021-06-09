package problem1486

func xorOperation(n int, start int) int {
	result := 0

	for i := 0; i < n ; i++{
		result = result ^ (start + 2 * i)
	}

	return result
}
