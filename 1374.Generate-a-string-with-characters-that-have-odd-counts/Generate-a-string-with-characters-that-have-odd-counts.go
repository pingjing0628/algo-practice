package problem1374

func generateTheString(n int) string {
	// Solution 1
	//     common := strings.Repeat("a", n - 1)

	//     if n % 2 != 0 {
	//         return "a" + common
	//     }


	//     return common + "b"

	// Solution 2
	result := make([]byte, n)
	for i := 0; i < n - 1; i++{
		result[i] = 'a'
	}
	if n % 2 != 0{
		result[n - 1] = 'a'
		return string(result)
	}

	result[ n - 1] = 'b'
	return string(result)
}
