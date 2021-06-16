package problem1732

func largestAltitude(gain []int) int {
	altitude := make([]int, len(gain) + 1)
	max := 0

	for i := 0; i < len(gain); i++{
		altitude[i + 1] = gain[i] + altitude[i]
		if altitude[i + 1] > max {
			max = altitude[i + 1]
		}
	}

	return max
}
