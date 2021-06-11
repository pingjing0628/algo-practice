package problem1688

func numberOfMatches(n int) int {
	matches := 0
	for n > 1 {
		result := n / 2
		matches += result
		n = n - result
	}

	return matches
}
