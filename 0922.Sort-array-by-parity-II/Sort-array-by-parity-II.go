package problem0922

func sortArrayByParityII(nums []int) []int {
	// Solution 1
	even, odd := 0, 1
	for even < len(nums) && odd < len(nums) {
		if nums[even] % 2 == 1 {
			nums[even], nums[odd] = nums[odd], nums[even]
			odd += 2
			continue
		}
		even += 2
	}
	return nums

	// Solution 2
	//result := make([]int, len(nums))
	//even, odd := 0, 1
	//for i := 0; i < len(nums); i++ {
	//	if nums[i] % 2 == 0{
	//		result[even] = nums[i]
	//		even += 2
	//	}else{
	//		result[odd] = nums[i]
	//		odd += 2
	//	}
	//}
	//
	//return result
}
