package problem1844

func replaceDigits(s string) string {
	result := []byte(s)
	for i := 1; i < len(s); i += 2 {
		result[i] = result[i - 1] + result[i] - '0'

	}
	return string(result)
}


