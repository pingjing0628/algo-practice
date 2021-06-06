package problem1528

func restoreString(s string, indices []int) string {
	result := make([]byte, len(s))
	for i, value := range indices{
		result[value] = s[i]
	}

	return string(result)
}
