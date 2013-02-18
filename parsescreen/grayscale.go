package main

import (
    "fmt"
    "os"
    "image/color"
    "image/png"
    "math"
    "code.google.com/p/probab/dst"
)

type ImageSet interface {
	Set(x, y int, c color.Color)
}

func main() {
	inname := "/Users/jpirruccello/Pictures/iPhoto Library/Masters/2013/02/17/20130217-110524/IMG_2018.PNG"
    outname := "/Users/jpirruccello/Desktop/New.png"

    file, err := os.Open(inname)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

    pic, err := png.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", inname, err)
		return
	}

    b := pic.Bounds()

	// Get an interface which can set pixels
	picSet := pic.(ImageSet)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			col := pic.At(x, y)
			r, g, bl, a := col.RGBA()
            
            
            intensity := gaussian2d(float64(x), float64(y), float64(b.Max.X), float64(b.Max.Y))
            r = uint32(intensity * float64(r))
            
            //Draw an X
            if math.Trunc(1000.0 * float64(x)/float64(b.Max.X)) == math.Trunc(1000.0 * float64(y)/float64(b.Max.Y)) || 
                    math.Trunc( 1000.0 * (1.0 - float64(x)/float64(b.Max.X))) == math.Trunc(1000.0 * float64(y)/float64(b.Max.Y)) {
                r, g, bl, a = 0, 0, 0, 0
            }

            //Luminosity method via http://www.johndcook.com/blog/2009/08/24/algorithms-convert-color-grayscale/
            avg := (21 * r + 71 * g + 7 * bl) / 100
            r, g, bl = avg, avg, avg

			// Update the colors
			newCol := color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(bl >> 8), uint8(a >> 8)}
			picSet.Set(x, y, newCol)
		}
	}

    // Write to the output file
	fd, err := os.Create(outname)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = png.Encode(fd, pic)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = fd.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func gaussian2d(x, y, xMax, yMax float64) float64 {
    result := 1 - math.Sqrt(math.Pow(x/xMax - 0.5, 2.0) + math.Pow(y/yMax - 0.5, 2.0))
    result = dst.NormalQtlFor(0.0, 3.0, result)
    if result < 0 {
        result = 0
    }
    //fmt.Println(x, x/xMax - 0.5, math.Pow(x/xMax - 0.5, 2.0), result, dst.NormalQtlFor(0.0, 1.0, result))
    if result > 1 {
        //fmt.Println(result)
        result = 1
    }

    return result
}