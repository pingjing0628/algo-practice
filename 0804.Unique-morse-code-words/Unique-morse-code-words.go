package problem0804

func uniqueMorseRepresentations(words []string) int {
	morse := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

	compare := make(map[string]int)

	for _, value := range words {
		wording := ""

		for _, val := range value {

			// val - 'a' 可取得 morse index
			// 透過 += 將字串組起來
			wording += morse[val - 'a']
		}
		// 組完成後將 compare 中此 wording 做累計
		// ex: map[".--...-.-." :1]
		compare[wording]++
	}

	return len(compare)
}
