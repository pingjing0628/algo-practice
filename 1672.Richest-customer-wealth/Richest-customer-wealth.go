package problem1672
func maximumWealth(accounts [][]int) int {
	max := 0
	for _, value := range accounts{
		sum := 0
		for _, val := range value {
			sum += val
			if max < sum {
				max = sum
			}
		}
	}
	return max
}