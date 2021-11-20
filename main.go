package main

import (
	"bytes"
	"fmt"
	"image"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No link provided")
		return
	}
	link := os.Args[1]

	var png []byte
	png, err := qrcode.Encode(link, qrcode.Highest, 0)
	if err != nil {
		panic(err)
	}
	imgconf, _, err := image.DecodeConfig(bytes.NewReader(png))
	img, _, err := image.Decode(bytes.NewReader(png))
	if err != nil {
		panic(err)
	}
	firstx := 9999
	firsty := 9999
	lastx := 0
	lasty := 0
	for y := 0; y <= imgconf.Height; y++ {
		for x := 0; x <= imgconf.Width; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			if r <= 0 {
				firstx = Min(firstx, x)
				firsty = Min(firsty, y)
			}
			if r == 0 {
				lastx = Max(lastx, x)
				lasty = Max(lasty, y)
			}
		}
	}

	for y := 0; y <= imgconf.Height; y++ {
		if (y < firsty-1) || (y > lasty+1) {
			continue
		}
		for x := 0; x <= imgconf.Width; x++ {
			if (x < firstx-1) || (x > lastx+1) {
				continue
			}
			r, _, _, _ := img.At(x, y).RGBA()
			if r > 0 {
				fmt.Print("\033[1;47m  \033[1;0m")
			} else {
				fmt.Print("\033[1;40m  \033[1;0m")
			}
		}
		fmt.Println()
	}
}
