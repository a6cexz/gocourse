package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	green := color.RGBA{0, 255, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	var red color.Color = color.RGBA{200, 30, 30, 255}

	y := 0
	for i := 0; i < 5; i++ {
		for x := 0; x < 200; x++ {
			rectImg.Set(x, y, red)
		}
		y += 40
	}

	x := 0
	for i := 0; i < 5; i++ {
		for y := 0; y < 200; y++ {
			rectImg.Set(x, y, red)
		}
		x += 40
	}

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}
