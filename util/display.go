package util

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/foobaz/lossypng/lossypng"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
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

func Img2CompressedPng(img image.Image, quality int) []byte {
	resized := imaging.Resize(img, W, 0, imaging.Lanczos)
	buffer := new(bytes.Buffer)
	if quality > 0 {
		compressed := lossypng.Compress(resized, lossypng.RGBAConversion, quality)
		png.Encode(buffer, compressed)
	} else {
		png.Encode(buffer, resized)
	}
	return buffer.Bytes()
}
