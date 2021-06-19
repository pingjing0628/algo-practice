package problem1309

import "strconv"

func freqAlphabets(s string) string {
	// s := "A1" // 分配儲存"A1"的内存空間，s 結構體裡的str指針指向這塊内存
	// s = "A2"  // 重新给"A2"的分配内存空間，s 結構體裡的str指針指向這塊内存

	// s := []byte{1} // 分配儲存1數組的内存空間，s結構體的array指針指向這個數組。
	// s = []byte{2}  // 將array的内容改為2
	// []byte和string的差別是更改變量的時候array的內容可以被更改。
	// string的指針指向的內容是不可以更改的，所以每更改一次字符串，就得重新分配一次內存

	result := []byte{}
	i := len(s) - 1
	for i >= 0{
		if s[i] == '#'{
			num, _ := strconv.Atoi(s[i - 2:i])
			result = append([]byte{byte(num) - 1 + 'a'}, result...)
			i -= 3
		}else{
			num := s[i] - '0'
			result = append([]byte{num - 1 + 'a'}, result...)
			i--
		}
	}
	return string(result)

	// string可以直接比較，而[]byte不可以，所以[]byte不可以當map的key值。
	// 因為無法修改string中的某個字符，需要粒度小到操作一個字符時，用[]byte。
	// string值不可為nil，所以如果你想要通過返回nil表達額外的含義，就用[]byte。
	// []byte切片這麼靈活，想要用切片的特性就用[]byte。
	// 需要大量字符串處理的時候用[]byte，性能好很多。
}
