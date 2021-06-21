package problem0728

func selfDividingNumbers(left int, right int) []int {
	res := []int{}

	for left <= right {
		v := left

		tmp := v
		for v > 0 {
			d := v % 10
			// 若為 10, 20...或為不能被整除的值則 break
			if d == 0 || (tmp % d) != 0 {
				break
			}
			v /= 10
		}

		if v == 0 {
			res = append(res, tmp)
		}
		left++
	}

	return res
}
