package problem0832

func flipAndInvertImage(image [][]int) [][]int {
	var result [][]int
	temp := []int{}

	for i := 0; i < len(image); i++{
		// 倒著 append 進入 temp
		for j := len(image) - 1; j >= 0; j-- {
			inv := invert(image[i][j])
			temp = append(temp, inv)
		}
		result = append(result, temp)
		temp = []int{}
	}

	return result

}

func invert(n int) int {
	if n == 0{
		return 1
	}else{
		return 0
	}
}