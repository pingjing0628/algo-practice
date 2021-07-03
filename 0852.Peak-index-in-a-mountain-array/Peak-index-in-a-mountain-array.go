package problem0852

func peakIndexInMountainArray(arr []int) int {
	// Solution 1
	//     for i := 0; i< len(arr); i++{
	//         if i == 0{
	//             continue
	//         }else if arr[i] > arr[i - 1] && arr[i] > arr[i + 1] {
	//             return i
	//         }
	//     }

	//     return 0

	// Solution 2
	i := 0
	for arr[i] < arr[i + 1] {
		i++
	}
	return i
}
