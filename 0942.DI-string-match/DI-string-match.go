package problem0942

func diStringMatch(s string) []int {
	min, max := 0, len(s)
	result := make([]int, len(s) + 1)
	for i := 0; i < len(s); i++{
		if s[i] == 'I'{
			result[i] = min
			min++
		}else if s[i] == 'D' {
			result[i] = max
			max--
		}
	}

	result[len(s)] = min

	return result
}
