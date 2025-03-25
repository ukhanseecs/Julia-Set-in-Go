package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"math/cmplx"
)

type Pixel struct {
	x float64
	y float64
}

// Function to check if point is in julia set
func isInJuliaSet(c complex128, pixel Pixel, maxIters int, scale float64) bool {
	z := complex(float64(pixel.x) / scale, float64(pixel.y) / scale)
	const escapeRadius = 2
	
	for i := 0; i < maxIters; i++ {
		z = z*z + c
		if cmplx.Abs(z) > escapeRadius {
			return false
		}
	}
	return true
}


func generateJuliaSetImage(c complex128, width int, height int, maxIters int, scale float64, filename string){
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	centerX := float64(width / 2)
	centerY := float64(height / 2)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pixel := Pixel{
				float64(x) - centerX, 
				float64(y) - centerY,
			}
			if isInJuliaSet(c, pixel, maxIters, scale) {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}
	file, err := os.Create(filename)
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err!=nil {
		fmt.Println("Error encoding image", err)
		return
	}

	fmt.Println("Image generated successfully")
}


func main() {
	
	image_height := 800
	image_width := 800
	maxIters := 100
	scale := 300.0
	juliaSetConstant := complex(-0.7, 0.27015)

	generateJuliaSetImage(juliaSetConstant, image_width, image_height, maxIters, scale, "julia.png")

}