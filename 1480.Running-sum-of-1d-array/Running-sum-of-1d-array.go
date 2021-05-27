package problem1480

func runningSum(nums []int) []int {
	//     output := make([]int, len(nums))
	//     for i := 0; i < len(nums); i++{
	//         detailSum := 0
	//         if i == 0 {
	//             output[i] = nums[i]
	//         }else {
	//             j := i
	//             for j != 0{
	//                 detailSum += nums[j]
	//                 j--
	//             }
	//             output[i] = detailSum + nums[0]

	//         }
	//     }
	//     return output

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
