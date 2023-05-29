package util

import (
	"bytes"
	//"github.com/foobaz/lossypng/lossypng"
	"github.com/kbinani/screenshot"
	"image"
	"image/jpeg"
)

var W, H int = 480, 640

func AvaiDisplays() []int {
	var displays []int
	for i := 0; i < screenshot.NumActiveDisplays(); i++ {
		displays = append(displays, i)
	}
	return displays
}

func Capture(display int) image.Image {
	bounds := screenshot.GetDisplayBounds(display)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	//buffer := new(bytes.Buffer)
	//png.Encode(buffer, img)
	//return buffer.Bytes()
	return img
}

func Image2Jpeg(img image.Image, quality int) []byte {
	buff := bytes.NewBuffer([]byte{})
	jpeg.Encode(buff, img, &jpeg.Options{Quality: quality})
	return buff.Bytes()
}
