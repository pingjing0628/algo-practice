package problem0682

import "strconv"

func calPoints(ops []string) int {
	toSum := make([]int, len(ops))

	for i := 0; i < len(ops); i++{
		if ops[i] == "C" {
			toSum = toSum[:len(toSum) - 1]
		} else if ops[i] == "D" {
			toSum = append(toSum, toSum[len(toSum) - 1] * 2)
		} else if ops[i] == "+"{
			toSum = append(toSum, toSum[len(toSum) - 1] + toSum[len(toSum) - 2])
		} else {
			num, _ := strconv.Atoi(ops[i])
			toSum = append(toSum, num)
		}
	}

	sum := Sum(toSum)
	return sum
}

func Sum(toSum []int) int {
	result := 0
	for _, value := range toSum {
		result += value
	}
	return result
}
