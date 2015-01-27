/*
From https://github.com/mjibson/go-dsp
*/
package main

import (
	//"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"

	"github.com/mjibson/go-dsp/fft"
)

func main() {
	original := make([]float64, 1020)
	Nitems := float64(len(original))

	i := float64(0)
	for k, _ := range original {
		original[k] = 255 * i / Nitems
		i++
	}

	transformed := fft.FFTReal(original)
	detransformed := fft.IFFT(transformed)
	
	fuxored := make([]complex128, len(original))
	copy(fuxored, transformed)
	for k, v := range fuxored {
		if cmplx.Abs(v) < 500 {
			fuxored[k] = 0 
		}
	}
	//fuxored[0] = 0
	defuxored := fft.IFFT(fuxored)

	//fmt.Println("Original ", original)
	//fmt.Println("Transformed ", transformed)
	//fmt.Println(fuxored)
	//fmt.Println("Detransformed ", detransformed)
	//fmt.Println("Defuxored ", defuxored)

	//Now plot shit
	Blocks := 6
	BlockHeight := 20
	img := image.NewRGBA(image.Rect(0, 0, len(original), Blocks * BlockHeight))
	block := -1
	
	block++
	for y := block * BlockHeight; y < (block+1) * BlockHeight; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			img.Set(x, y, color.RGBA{uint8(original[x]), 0, 0, 0xff})
		}
	}
	
	block++
	for y := block * BlockHeight; y < (block+1) * BlockHeight; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			img.Set(x, y, color.RGBA{uint8(cmplx.Abs(transformed[x])), 0, 0, 0xff})
		}
	}
	
	block++
	for y := block * BlockHeight; y < (block+1) * BlockHeight; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			img.Set(x, y, color.RGBA{uint8(cmplx.Abs(detransformed[x])), 0, 0, 0xff})
		}
	}
	
	block++
	for y := block * BlockHeight; y < (block+1) * BlockHeight; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			img.Set(x, y, color.RGBA{uint8(original[x]), 0, 0, 0xff})
		}
	}
	
	block++
	for y := block * BlockHeight; y < (block+1) * BlockHeight; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			img.Set(x, y, color.RGBA{uint8(cmplx.Abs(fuxored[x])), 0, 0, 0xff})
		}
	}
	
	block++
	for y := block * BlockHeight; y < (block+1) * BlockHeight; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			img.Set(x, y, color.RGBA{uint8(cmplx.Abs(defuxored[x])), 0, 0, 0xff})
		}
	}

	file, err := os.Create("simple.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	png.Encode(file, img)
}
