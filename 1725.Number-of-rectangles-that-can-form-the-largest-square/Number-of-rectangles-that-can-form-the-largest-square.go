package problem1725

func countGoodRectangles(rectangles [][]int) int {
	minmum := make([]int, 0)
	count := 0

	for _, index := range rectangles {

		if index[0] >= index[1] {
			minmum = append(minmum, index[1])
		}else {
			minmum = append(minmum, index[0])
		}
	}
	maximum := max(minmum)

	for i := 0; i < len(minmum); i++{
		if maximum == minmum[i] {
			count++
		}
	}

	return count

}

func max (input []int) int {
	max := 0
	for i := 0; i < len(input); i++{
		if max < input[i] {
			max = input[i]
		}
	}
	return max
}
