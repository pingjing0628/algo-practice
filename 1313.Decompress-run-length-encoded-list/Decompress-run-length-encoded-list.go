package problem1313

func decompressRLElist(nums []int) []int {
	compressList := make([]int, 0)

	for i := 0; i < len(nums); i += 2 {
		for j := 1; j <= nums[i]; j++ {
			compressList = append(compressList, nums[i + 1])
		}
	}

	return compressList
}