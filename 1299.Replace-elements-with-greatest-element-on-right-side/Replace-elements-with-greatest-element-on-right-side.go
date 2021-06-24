package problem1299

func replaceElements(arr []int) []int {
	// Solution 1
	temp := 0
	maxRight := arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		temp = arr[i]
		arr[i] = maxRight
		if temp > maxRight {
			maxRight = temp
		}
	}
	arr[len(arr)-1] = -1
	return arr

	// Solution 2
	//     result := make([]int, len(arr))
	//     maxi := 0
	//     for i := 0; i < len(arr); i++{
	//         maxi = max(arr[i+1:])
	//         result[i] = maxi

	//         if len(arr[i+1:]) == 0 {
	//             result[i] = -1
	//         }

	//     }

	//     return result
	// }

	// func max (input []int) int {
	//     maxi := 0
	//     for _, v := range input {
	//         if v > maxi {
	//             maxi = v
	//         }
	//     }

	//     return maxi
	// }
}
