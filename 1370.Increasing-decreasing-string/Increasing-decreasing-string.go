package problem1370

func sortString(s string) string {
	counts := make([]int, 26)
	stringLength := len(s)
	result := ""

	// make count in counts
	// [4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
	for _, value := range s {
		counts[value - 'a']++
	}

	for stringLength > 0 {
		// smallest
		for i := 0; i < 26; i++{
			if counts[i] > 0{
				result += string(i + 'a')
				counts[i]--
				stringLength--
			}
		}

		// largest
		for i := 25; i >= 0; i-- {
			if counts[i] > 0{
				result += string(i + 'a')
				counts[i]--
				stringLength--
			}
		}
	}

	return result
}
