package problem0709

func toLowerCase(s string) string {
	// Solution 1
	result := []byte(s)
	for i, value := range result{
		if value >= 'A' && value <= 'Z' {
			result[i] = value + 32
		}
	}
	return string(result)

	// Solution 2
	//return strings.ToLower(s)
}
