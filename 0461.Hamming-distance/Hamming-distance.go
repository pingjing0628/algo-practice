package problem0461

func hammingDistance(x int, y int) int {
	//     diff := x ^ y
	//     result := 0
	//     for diff > 0 {
	//         if diff & 1 == 1 {
	//             result++
	//         }

	//         diff >>= 1
	//     }

	//     return result

	count := 0
	for x > 0 || y > 0 {
		if x % 2 != y % 2 {
			count++
		}
		x /= 2
		y /= 2
	}
	return count
}
