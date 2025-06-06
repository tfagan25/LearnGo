package main

import (
	"fmt"
	"image"
	"os"
	_ "image/png"
	"image/color"
)

func loadImageFromFile(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    image, _, err := image.Decode(f)
    return image, err
}

func outputLetter(letterImage image.Image, size int) {
	bounds := letterImage.Bounds()
	cols, rows := bounds.Max.X / size, bounds.Max.Y / size
	white := color.RGBA{255, 255, 255, 255}
	r1, g1, b1, a1 := white.RGBA()

	for i := range rows {
		for j := range cols {
			r, g, b, a := letterImage.At(j * size, i * size).RGBA()
			if r == r1 && g == g1 && b == b1 && a == a1 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func main () {
	image, err := loadImageFromFile("letter-a.png")
	if err != nil {
		fmt.Println("Error loading image:", err)
	}

	outputLetter(image, 10)
}