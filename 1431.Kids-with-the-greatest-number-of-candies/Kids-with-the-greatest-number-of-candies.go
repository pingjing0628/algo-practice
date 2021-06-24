package problem1431

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	max := 0

	for _, value := range candies {
		if value > max {
			max = value
		}
	}

	for i, val := range candies {
		if val + extraCandies >= max {
			result[i] = true
		}else {
			result[i] = false
		}
	}

	return result
}
