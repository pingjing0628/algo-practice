package problem1773

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	match := map[string]int {
		"type": 0,
		"color": 1,
		"name": 2,
	}
	count := 0

	for i := range items{
		if items[i][match[ruleKey]] == ruleValue {
			count++
		}
	}

	return count

}
