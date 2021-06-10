package problem1662

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	// Solution 1
	//     w1 := ""
	//     w2 := ""
	//     for _, v := range word1 {
	//         w1 += v
	//     }

	//     for _, v := range word2 {
	//         w2 += v
	//     }

	//     if w1 == w2 {
	//         return true
	//     }

	// Solution 2
	if strings.Join(word1[:],"") == strings.Join(word2[:],"") {
		return true
	}

	return false
}
