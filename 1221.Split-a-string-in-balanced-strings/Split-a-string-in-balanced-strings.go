package problem1221

func balancedStringSplit(s string) int {
	match := 0
	count := 0

	for i, value := range s {
		if string(value) == "L" {
			match++
		}

		if string(value) == "R" {
			match--
		}

		if i != 0 && match == 0 {
			count++
		}
	}

	return count
}
