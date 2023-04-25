package util

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
)

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

func Img2CompressedPng(img image.Image) []byte {
	resized := imaging.Resize(img, 1045, 0, imaging.Lanczos)
	buffer := new(bytes.Buffer)
	png.Encode(buffer, resized)
	return buffer.Bytes()
}