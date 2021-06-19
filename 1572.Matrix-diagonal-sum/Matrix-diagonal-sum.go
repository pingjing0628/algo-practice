package problem1572

func diagonalSum(mat [][]int) int {
	matLen := len(mat)
	result := 0
	for i := 0; i< matLen; i++{
		result += mat[i][i] + mat[i][matLen - i - 1]
	}

	if matLen % 2 != 0{
		result -= mat[matLen / 2][matLen / 2]
	}

	return result
}
