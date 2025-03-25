package fractal

import "math/cmplx"

const escapeRadius = 2

type Pixel struct {
	X float64
	Y float64
}


// Function to check if point is in julia set
func IsInJuliaSet(c complex128, pixel Pixel, maxIters int, scale float64) bool {
	z := complex(float64(pixel.X) / scale, float64(pixel.Y) / scale)

	for i := 0; i < maxIters; i++ {
		z = z*z + c
		if cmplx.Abs(z) > escapeRadius {
			return false
		}
	}
	return true
}