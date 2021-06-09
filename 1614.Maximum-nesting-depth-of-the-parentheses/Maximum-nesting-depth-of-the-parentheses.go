package problem1614

func maxDepth(s string) int {
	max := 0
	result := 0

	for _, val := range s {
		if val == '(' {
			max++
			if max > result {
				result = max
			}
		} else if val == ')' {
			max--
		}
	}

	return result
}