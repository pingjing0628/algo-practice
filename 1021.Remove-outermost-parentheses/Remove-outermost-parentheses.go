package problem1021

import "bytes"

func removeOuterParentheses(s string) string {
	//     count := 0
	//     var result []uint8

	//     for a, i := 0, 0; i < len(s); i++{
	//         if s[i] == '(' {
	//             count++
	//         }else {
	//             count--
	//             if count == 0 {
	//                 result = append(result, s[a+1: i]...)
	//                 a = i + 1
	//             }
	//         }
	//     }

	//     return string(result)

	var depth int
	var res bytes.Buffer
	for i := range s {
		switch s[i] {
		case '(':
			if depth > 0 {
				res.WriteString(s[i:i+1])
			}
			depth++
		case ')':
			depth--
			if depth > 0 {
				res.WriteString(s[i:i+1])
			}
		}
	}
	return res.String()
}
