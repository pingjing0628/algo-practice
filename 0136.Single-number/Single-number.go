package problem0136

func singleNumber(nums []int) int {
	//     m:= make(map[int]bool,0)

	//     for _,n := range nums{
	//         if _,ok := m[n]; !ok{
	//             m[n] = true
	//         } else{
	//             m[n] = false
	//         }
	//     }

	//     for k,v := range m{
	//         if v {
	//             return k
	//         }
	//     }
	//     return -1

	result := 0
	for _, value := range nums{
		result = result ^ value
	}
	return result
}