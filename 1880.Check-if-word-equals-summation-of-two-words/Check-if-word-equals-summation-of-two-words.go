package problem1880

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return transfer(firstWord) + transfer(secondWord) == transfer(targetWord)
}

func transfer(s string) int {
	sum := 0
	for _, val := range s{
		sum = sum * 10 + int(val - 'a')
	}
	return sum
}
