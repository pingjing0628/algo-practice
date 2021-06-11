package problem1588

func sumOddLengthSubarrays(arr []int) int {
    sum := 0
    for i := 1; i <= len(arr); i += 2 {
        for j := 0; j <= len(arr) - i; j++{
            for k := j; k < j + i; k++ {
                sum += arr[k]
            }
        }
    }
    return sum
}

// OOP
//type Custom struct {
//	sum int
//	num int
//}
//
//func (this *Custom) add(n int) {
//	this.sum += n
//	this.num++
//}
//
//func (this *Custom) minus(n int) int {
//	this.sum -= n
//	this.num--
//	return this.sum
//}
//
//func (this *Custom) getSum() int {
//	return this.sum
//}
//
//func (this *Custom) getNum() int {
//	return this.num
//}
//
//func sumOddLengthSubarrays(arr []int) int {
//	numSub := len(arr)/2 + len(arr)%2
//	record := make([]*Custom, numSub)
//	for i:=0; i<numSub; i++ {
//		record[i] = &Custom {
//			sum: 0,
//			num: 0,
//		}
//	}
//	ret := 0
//
//	for idx, num := range arr {
//		for i := 0; i < numSub; i++ {
//			record[i].add(num)
//			if record[i].getNum() == i*2 + 1 {
//				ret += record[i].getSum()
//			} else if record[i].getNum() > i*2 + 1 {
//				ret += record[i].minus(arr[idx-(i*2+1)])
//			}
//		}
//	}
//	return ret
//}
