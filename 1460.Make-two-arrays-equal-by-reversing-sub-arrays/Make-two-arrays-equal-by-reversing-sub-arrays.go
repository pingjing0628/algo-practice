package problem1460

func canBeEqual(target []int, arr []int) bool {
	// Solution 1 : Runtime 12 ms, Memory 4.1 MB
	//  sort.Ints(target)
	// 	sort.Ints(arr)

	// 	for i := 0; i < len(target); i++ {
	// 		if target[i] != arr[i] {
	// 			return false
	// 		}
	// 	}
	// 	return true

	// Solution 2: Runtime 4 ms, Memory 5.3 MB
	mapping := make(map[int]int)

	for i := 0; i < len(target); i++{
		mapping[target[i]]++
		mapping[arr[i]]--
	}

	for _, value := range mapping {
		if value != 0 {
			return false
		}
	}

	return true
}
