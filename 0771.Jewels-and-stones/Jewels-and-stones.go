package problem0771

import "strings"

func numJewelsInStones(jewels string, stones string) int {
	//     count := 0
	//     for _, value := range jewels {

	//         for _, val := range stones {
	//             if value == val {
	//                 count++
	//             }
	//         }
	//     }

	//     return count

	count := 0
	for _, val:= range jewels {
		count += strings.Count(stones, string(val))
	}

	return count
}
