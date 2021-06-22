package problem0905

func sortArrayByParity(nums []int) []int {
	// Solution 1
	backLength := len(nums) - 1
	frontLength := 0
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++{
		if nums[i] % 2 != 0 {
			result[backLength] = nums[i]
			backLength--
		}else{
			result[frontLength] = nums[i]
			frontLength++
		}
	}
	return result

	// Solution 2
	//var n, m []int
	//for i := 0; i < len(nums); i++ {
	//	if nums[i] % 2 == 0 {
	//		m = append(m, nums[i])
	//	} else {
	//		n = append(n, nums[i])
	//	}
	//}
	//m = append(m, n...)
	//return m
}
