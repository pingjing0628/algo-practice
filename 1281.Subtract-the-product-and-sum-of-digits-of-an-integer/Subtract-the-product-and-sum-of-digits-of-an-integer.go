package problem1281

func subtractProductAndSum(n int) int {
	prod, sum := 1, 0

	for n > 0 {
		num := n % 10
		prod *= num
		sum += num
		n /= 10
	}

	return prod - sum
}