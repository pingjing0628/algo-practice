package problem1704

func halvesAreAlike(s string) bool {
	vowels := map[byte]bool {
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	countR, countL := 0, 0

	for i := 0; i < len(s) / 2; i++ {
		if vowels[s[i]] {
			countL++
		}
	}

	for i := len(s)/2; i < len(s); i++{
		if vowels[s[i]] {
			countR++
		}
	}
	return countL == countR
}
