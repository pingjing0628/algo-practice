package problem1337

import "sort"

// Solution 1
type entry struct{
	index int
	soldiers int
}

func kWeakestRows(mat [][]int, k int) []int {
	res := make([]int, 0, k)
	count := make([]entry, len(mat), len(mat))
	// 組成 count[index]
	// ex: [{0 2} {1 4} {2 1} {3 2} {4 5}]
	for r, row := range mat {
		soldiers := 0
		for _, val := range row {
			soldiers += val
		}
		count[r] = entry{r, soldiers}
	}

	// 若 i, j soldiers 相同, 則 i 的 index 須小於 j 的 index
	// 或 j 的 soldiers 需大於 i
	sort.Slice(count, func(i, j int) bool {
		si, sj := count[i].soldiers, count[j].soldiers
		return (si == sj && count[i].index < count[j].index) || si < sj
	})
	// 根據 soldiers value sort 後
	// [{2 1} {0 2} {3 2} {1 4} {4 5}]

	// 將排序好的 count 加入到回傳值
	for i := 0; i < k; i++ {
		res = append(res, count[i].index)
	}

	return res
}

// Solution 2
//func kWeakestRows(mat [][]int, k int) []int {
//	deleteRow := make(map[int]bool)
//	m, n := len(mat), len(mat[0])
//	ans := []int{}
//	for j := 0; j < n; j++ {
//		for i := 0; i < m; i++ {
//			if deleteRow[i] || mat[i][j] == 1 { continue }
//			deleteRow[i] = true
//			ans = append(ans, i)
//			if len(ans) == k { return ans }
//		}
//	}
//	if len(ans) < k {
//		for i := 0; i < m; i++ {
//			if deleteRow[i] { continue }
//			ans = append(ans, i)
//			if len(ans) == k { return ans }
//		}
//	}
//	return ans
//}
