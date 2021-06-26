package problem1742

func countBalls(lowLimit int, highLimit int) int {
	result := 0
	mapping := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++{
		j := i
		sum := 0
		for j > 0 {
			num := j % 10
			sum += num
			j /= 10
		}

		if _, exist := mapping[sum]; exist {
			mapping[sum]++
		} else {
			mapping[sum] = 1
		}

		result = max(result, mapping[sum])
	}

	return result
}

func max(num, result int) int {
	if num > result {
		return num
	}
	return result
}
