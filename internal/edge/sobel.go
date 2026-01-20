package edge

import (
	"image"
	"math"
)

type Edge struct {
	Strength  float64
	Direction float64
}

func Sobel(img image.Image) [][]Edge {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Sobel kernels
	gx := [3][3]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
	gy := [3][3]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	edges := make([][]Edge, height)
	for y := 0; y < height; y++ {
		edges[y] = make([]Edge, width)
		for x := 0; x < width; x++ {
			var sumGx, sumGy float64

			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					px := x + kx
					py := y + ky

					if px < 0 {
						px = 0
					} else if px >= width {
						px = width - 1
					}
					if py < 0 {
						py = 0
					} else if py >= height {
						py = height - 1
					}

					c := img.At(bounds.Min.X+px, bounds.Min.Y+py)
					r, g, b, _ := c.RGBA()
					gray := float64(r>>8)*0.2126 + float64(g>>8)*0.7152 + float64(b>>8)*0.0722

					sumGx += gray * float64(gx[ky+1][kx+1])
					sumGy += gray * float64(gy[ky+1][kx+1])
				}
			}

			magnitude := math.Sqrt(sumGx*sumGx + sumGy*sumGy)
			direction := math.Atan2(sumGy, sumGx)

			edges[y][x] = Edge{
				Strength:  magnitude,
				Direction: direction,
			}
		}
	}

	return edges
}
