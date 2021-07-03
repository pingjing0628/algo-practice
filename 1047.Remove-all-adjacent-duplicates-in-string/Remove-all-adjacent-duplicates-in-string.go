package problem1047

func removeDuplicates(s string) string {
	words := []byte{}

	for i := 0; i< len(s); i++{
		// 若相同則不取重複的值
		if len(words) != 0 && words[len(words) - 1] == s[i] {
			words = words[:len(words) - 1]
		}else{
			words = append(words, s[i])

		}
	}

	return string(words)
}
