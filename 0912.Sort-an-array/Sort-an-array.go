package problem0912

import "fmt"

func sortArray(nums []int) []int {
	fmt.Println(nums)
	nums = quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(nums []int, left, right int) []int {
	if left > right {
		return nums
	}

	pivot := nums[left]
	i := left
	j := right

	for i != j {
		for nums[j] >= pivot && i < j {
			j--
		}
		for nums[i] <= pivot && i < j {
			i++
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[left], nums[i] = nums[i], nums[left]
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)

	return nums
}
