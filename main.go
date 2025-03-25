package main

import (
	"julia/imagegen"
)


func main() {
	
	image_height := 800
	image_width := 800
	maxIters := 100
	scale := 300.0
	juliaSetConstant := complex(-0.7, 0.27015)

	imagegen.GenerateJuliaSetImage(juliaSetConstant, image_width, image_height, maxIters, scale, "julia.png")

}