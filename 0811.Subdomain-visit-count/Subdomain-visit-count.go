package problem0811

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	counting := make(map[string]int)
	result := []string{}

	for _, value := range cpdomains {
		// 根據空格切割字串
		parts := strings.Split(value, " ")

		if len(parts) >= 2 {
			// 將字串第一個轉為數字, 以作運算
			times, _ := strconv.Atoi(parts[0])
			// 將剩餘字串根據 . 做分割
			domain := strings.Split(parts[1],".")

			// 將剩餘字串根據域名由頭至尾依序加入至 map 中, 並以數量 times 作為 value
			// 若有重複則相加
			for i := 0; i < len(domain); i++{
				v := strings.Join(domain[i:len(domain)], ".")
				counting[v] = counting[v] + times
			}
		}

	}

	// 將字串組成回傳格式
	for i, val := range counting {
		// 將 map 中數字轉回字串
		allDomain := strconv.Itoa(val) + " " + i
		result = append(result, allDomain)
	}

	return result
}
