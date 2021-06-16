package problem1266

func minTimeToVisitAllPoints(points [][]int) int {
	result := 0
	Max := func (val1, val2 int) int {
		if val1 > val2 {
			return val1
		}

		return val2
	}

	Abs := func (val int) int {
		if val < 0 {
			return -val
		}

		return val
	}

	for i := 1; i < len(points); i++ {
		result += Max(Abs(points[i - 1][0] - points[i][0]), Abs(points[i - 1][1] - points[i][1]))
	}
	return result
}
