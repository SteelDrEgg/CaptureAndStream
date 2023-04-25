package util

import (
	"strconv"
)

var warehouse = new([100][]byte)
var curIndex = 0

func AddData(data []byte) {
	if curIndex == 100 {
		curIndex = 0
	}
	warehouse[curIndex] = data
	curIndex++
}

func ReadData(index string) []byte {
	//strIndex := fmt.Sprintf("%d", index)
	//if len(index) > 2 {
	//	index = index[len(index)-3:]
	if len(index) > 1 {
		index = index[len(index)-2:]
	} else {
		index = index[len(index)-1:]
	}
	rIndex, _ := strconv.Atoi(index)
	return warehouse[rIndex]
}
