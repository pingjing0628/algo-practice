package problem1450

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	count := 0
	for i := 0; i < len(startTime); i++{
		if startTime[i] <= queryTime && queryTime <= endTime[i]{
			count++
		}

	}

	return count
}
