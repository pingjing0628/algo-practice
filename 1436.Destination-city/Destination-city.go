package problem1436

func destCity(paths [][]string) string {
	countPaths := make(map[string]int)

	// First to count the appearance of city
	// i < 2, 2 means only has the origin and destination two value in 2-D array
	for _, value := range paths {
		for i := 0; i < 2; i++{
			countPaths[value[i]]++
		}
	}

	// Second to get the value in countPaths which is 1 and is the destination i == 1
	for _, val := range paths {
		for i := 0; i < 2; i++{
			if countPaths[val[i]] == 1 && i == 1 {
				return val[i]
			}
		}
	}

	return ""
}