package problem1748

func sumOfUnique(nums []int) int {
	// Solution 1
	//     count := map[int]int{}
	//     sum := 0
	//     for i := 0; i < len(nums); i++{
	//         _, exist := count[nums[i]]
	//         if exist {
	//             count[nums[i]]++
	//         }else{
	//             count[nums[i]] = 1
	//         }
	//     }
	//     for key, value := range count{
	//         if value == 1{
	//             sum += key
	//         }
	//     }

	//     return sum

	// Solution 2
	ans, tmp := 0, make([]int, 101)
	for _, n := range nums {
		if tmp[n] == 0 {
			tmp[n] = 1
		} else {
			tmp[n] = -1
		}
	}
	for i, v := range tmp {
		if v == 1 { ans += i }
	}
	return ans
}
