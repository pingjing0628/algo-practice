package problem0657

func judgeCircle(moves string) bool {
	sum1, sum2 := 0, 0

	for _, value := range moves {
		switch value {
		case 'U':
			sum1++
		case 'D':
			sum1--
		case 'L':
			sum2--
		case 'R':
			sum2++
		}
	}

	if sum1 == 0 && sum2 == 0 {
		return true
	}
	return false
}
