package problem0961

func repeatedNTimes(nums []int) int {
	counts := map[int]int{}

	for _, a := range nums {
		if _, ok := counts[a]; ok {
			return a
		}
		counts[a] = 1
	}

	return 0
}
