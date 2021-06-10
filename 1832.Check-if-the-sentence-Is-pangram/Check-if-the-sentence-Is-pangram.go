package problem1832

func checkIfPangram(sentence string) bool {
	// Solution 1
	//     var result[26]int
	//     for i := 0; i < len(sentence); i++{
	//         if sentence[i] >= 'a' && sentence[i] <= 'z' {
	//             result[sentence[i] - 'a'] = 1
	//         }
	//     }

	//     for i := 0; i < len(result); i++{
	//         if result[i] == 0 {
	//             return false
	//         }
	//     }

	//     return true

	// Solution 2
	seen := map[rune]bool{}

	for _, letter := range sentence {
		seen[letter] = true
	}

	return len(seen) == 26
}
