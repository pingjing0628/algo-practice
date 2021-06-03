package problem1370

func sortString(s string) string {
	counts := make([]int, 26)
	stringLength := len(s)
	result := ""
	for _, value := range s {
		counts[value - 'a']++
	}

	for stringLength > 0 {
		for i := 0; i < 26; i++{
			if counts[i] > 0{
				result += string(i + 'a')
				counts[i]--
				stringLength--
			}
		}

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
