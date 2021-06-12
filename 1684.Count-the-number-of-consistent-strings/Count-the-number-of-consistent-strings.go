package problem1684

func countConsistentStrings(allowed string, words []string) int {
	result := len(words)
	alph := [26]int{}

	for _, value := range allowed {
		alph[value - 'a'] = 1
	}

	for _, vals := range words {
		for _, word_val := range vals {
			if alph[word_val - 'a'] == 0{
				result--
				break
			}
		}
	}

	return result
}
