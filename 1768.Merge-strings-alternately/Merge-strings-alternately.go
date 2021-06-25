package problem1768

func mergeAlternately(word1 string, word2 string) string {
	result := []byte{}
	one, two := []byte(word1), []byte(word2)

	for i, j := 0, 0; i < len(word1) || j < len(word2); i, j = i+1, j+1{
		//
		if i == len(word1) {
			return string(append(result, two[j:]...))
		}
		if j == len(word2) {
			return string(append(result, one[i:]...))
		}
		result = append(result, []byte{ one[i], two[j] }...)

		// append a string to a byte slice:
		// slice := append([]byte("Hello "), "world!"...)
	}

	return string(result)
}
