package problem1816

func truncateSentence(s string, k int) string {
	// Solution 1
	//     splitS := strings.Split(s, " ")

	//     return strings.Join(splitS[:k], " ")

	// Solution 2
	result := []byte{}
	for _, val := range []byte(s) {
		if val == ' ' {
			k--
		}
		if k == 0 {
			break
		}
		result = append(result, val)
	}
	return string(result)
}
