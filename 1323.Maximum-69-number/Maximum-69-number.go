package problem1323

import (
	"strconv"
	"strings"
)

func maximum69Number (num int) int {
	// Solution1
	//     result, times := 0, 1
	//     nums := make([]int, 0)

	//     for num % 10 != 0{
	//         nums = append(nums, num % 10)
	//         num = num / 10

	//     }

	//     for i := len(nums) - 1; i >= 0; i-- {
	//         if nums[i] != 9 {
	//             nums[i] = 9
	//             break
	//         }
	//     }
	//     for i := 0; i < len(nums); i++{
	//         result += nums[i] * times
	//         times *= 10
	//     }

	//     return result

	// Solution 2
	// Itoa: int to string
	// Atoi: string to int
	result, _ := strconv.Atoi(strings.Replace(strconv.Itoa(num), "6", "9", 1))
	return result
}
