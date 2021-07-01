package problem1207

func uniqueOccurrences(arr []int) bool {
	mapping := make(map[int]int)
	compare := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		mapping[arr[i]]++
	}

	for _, value := range mapping{
		compare[value]++
		if compare[value] > 1 {
			return false
		}
	}

	return true
}
