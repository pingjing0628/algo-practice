package problem1351

func countNegatives(grid [][]int) int {
	count := 0
	for _, array := range grid {
		for _, value := range array {
			if value < 0 {
				count++
			}
		}
	}
	return count
}