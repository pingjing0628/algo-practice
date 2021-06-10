package problem1656

type OrderedStream struct {
	ptr int
	stringVal []string
}


func Constructor(n int) OrderedStream {
	return OrderedStream{0, make([]string, n + 1)}
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
	result := []string{}

	this.stringVal[idKey - 1] = value // 將 value 根據 idKey 放至 stringVal 陣列中

	// 若 stringVal key ptr 位置不為空, 則將此位置值放至 result 中, 並繼續找尋下一個值
	for this.stringVal[this.ptr] != "" {
		result = append(result, this.stringVal[this.ptr])
		this.ptr++
	}

	return result
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
