package problem1812

func squareIsWhite(coordinates string) bool {
	// Solution 1
	//     blackMap := map[string]bool{
	//         "a": true,
	//         "c": true,
	//         "e": true,
	//         "g": true,
	//     }

	//     v, _ := strconv.Atoi(coordinates[1:])
	//     _, ok := blackMap[coordinates[:1]]
	//     if v % 2 != 0 && ok{
	//         return false
	//     }
	//     if v % 2 == 0 && !ok{
	//         return false
	//     }

	//     return true

	// Solution 2
	return !(int(coordinates[1] - '0') % 2 == (int(coordinates[0] - 'a') + 1) % 2)
}
