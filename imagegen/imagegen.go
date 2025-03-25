package imagegen

import(
	"image"
	"image/color"
	"julia/fractal"
	"julia/fileutil"
)


func GenerateJuliaSetImage(c complex128, width int, height int, maxIters int, scale float64, filename string){
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	centerX := float64(width / 2)
	centerY := float64(height / 2)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pixel := fractal.Pixel{
				X: float64(x) - centerX, 
				Y: float64(y) - centerY,
			}
			if fractal.IsInJuliaSet(c, pixel, maxIters, scale) {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}
	fileutil.SaveImage(img, filename)
}



