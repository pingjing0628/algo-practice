package problem1678

import "bytes"

func interpret(command string) string {
	// Solution 1
	//     result := ""
	//     for i := 0; i < len(command); i++ {
	//         switch{
	//             case command[i] == 'G':
	//                 result += "G"
	//             case command[i:i+2] == "()":
	//                 result += "o"
	//                 i += 1
	//             default:
	//                 result += "al"
	//                 i += 3
	//         }

	//     }

	//     return result

	// Solution 2
	var buffer bytes.Buffer // 使用 bytes.Buffer 來組裝 string
	temp := []rune {}
	dict := map[string]string {
		"G": "G",
		"()": "o",
		"(al)": "al",
	}
	for _, v := range command {
		temp = append(temp, v)
		if val, ok := dict[string(temp)]; ok {
			buffer.WriteString(val) // 將 string 寫入到 buffer 尾部
			temp = []rune {}
		}
	}
	return buffer.String() //ReadString

	// Go中可以使用 "+" 合併 string，但是這種合併方式效率非常低，每合併一次，都是創建一個新的 string,就必須遍歷複製一次 string
}
