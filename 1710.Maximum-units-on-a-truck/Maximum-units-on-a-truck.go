package problem1710

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sum := 0
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	for _, box := range boxTypes {
		if box[0] <= truckSize {
			sum += box[0] * box[1]
			truckSize -= box[0]
		}else {
			sum += box[1] * truckSize
			break
		}
	}

	return sum
}
