package problem1636

import (
	"fmt"
	"sort"
)

func frequencySort(nums []int) []int {
	countMap := map[int]int{}

	for _, value := range nums{
		countMap[value]++
	}
	//for i := 0 ; i < len(nums); i++ {
	//	count := 1
	//
	//	_, ok := countMap[nums[i]]
	//	if ok {
	//		countMap[nums[i]] += 1
	//	} else {
	//		countMap[nums[i]] = count
	//	}
	//
	//}

	sort.Slice(nums, func(i, j int) bool{
		// 若 map 中 所 count 數量相同, 則回傳 值大的
		if countMap[nums[i]] == countMap[nums[j]] {
			return nums[j] < nums[i]
		}

		// 回傳 count 數量大的值
		return countMap[nums[i]] < countMap[nums[j]]
	})


	return nums
}

func Go()  {
	nums := []int{2, 3, 1, 3, 2}
	//nums := []int{1, 1, 2, 2, 2, 3}
	nums = frequencySort(nums)

	fmt.Println(nums)
}
