package util

import (
	"os"
)

func Byte2png(img []byte, fileName string) {
	file, _ := os.Create("./temp/" + fileName + ".png")
	file.Write(img)
	defer file.Close()
}
