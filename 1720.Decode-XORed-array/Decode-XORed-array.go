package problem1720

func decode(encoded []int, first int) []int {
	arr := make([]int, 0)
	arr = append(arr, first)

	for i := 0; i < len(encoded); i++{
		first = encoded[i] ^ first
		arr = append(arr, first)
	}

	return arr
}
