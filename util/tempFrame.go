package util

import (
	"strconv"
)

var warehouse = new([30][]byte)
var curIndex = 0

func AddData(data []byte) {
	if curIndex == 10 {
		curIndex = 0
	}
	warehouse[curIndex] = data
	curIndex++
}

func ReadData(index string) []byte {
	//strIndex := fmt.Sprintf("%d", index)
	index = index[len(index)-1:]
	rIndex, _ := strconv.Atoi(index)
	return warehouse[rIndex]
}