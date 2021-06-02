package problem0344

func reverseString(s []byte) []byte {
	//length := len(s) - 1
	//
	//for i := 0; i <= length/2; i++ {
	//	s[i], s[length] = s[length], s[i]
	//	length--
	//}

	left := 0
	right := len(s) - 1
	for right > left {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}

	return s
}
