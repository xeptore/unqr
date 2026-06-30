package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func main() {
	os.Exit(run())
}

func run() (code int) {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <image>\n", os.Args[0])
		return 2
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "open image: %v\n", err)
		return 1
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "close image: %v\n", err)
			if code == 0 {
				code = 1
			}
		}
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decode image: %v\n", err)
		return 1
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "prepare bitmap: %v\n", err)
		return 1
	}

	result, err := qrcode.NewQRCodeReader().Decode(bmp, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decode QR: %v\n", err)
		return 1
	}

	fmt.Println(result.GetText())
	return 0
}
