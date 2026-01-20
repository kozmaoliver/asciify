package imageio

import (
	"image"
	"image/color"
	"math"
)

// DifferenceOfGaussians applies a Difference of Gaussians filter to an image.
func DifferenceOfGaussians(img image.Image, sigma1, sigma2 float64) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	blurred1 := gaussianBlur(img, sigma1)
	blurred2 := gaussianBlur(img, sigma2)

	result := image.NewGray(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c1 := blurred1.At(bounds.Min.X+x, bounds.Min.Y+y)
			c2 := blurred2.At(bounds.Min.X+x, bounds.Min.Y+y)

			r1, g1, b1, _ := c1.RGBA()
			r2, g2, b2, _ := c2.RGBA()

			gray1 := float64(r1>>8)*0.2126 + float64(g1>>8)*0.7152 + float64(b1>>8)*0.0722
			gray2 := float64(r2>>8)*0.2126 + float64(g2>>8)*0.7152 + float64(b2>>8)*0.0722

			// Normalize
			diff := gray1 - gray2
			value := (diff + 255.0) / 2.0
			
			// Clamp to valid range
			if value > 255.0 {
				value = 255.0
			}
			if value < 0.0 {
				value = 0.0
			}

			result.SetGray(x, y, color.Gray{Y: uint8(value)})
		}
	}

	return result
}

func gaussianBlur(img image.Image, sigma float64) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	kernelSize := int(math.Ceil(3*sigma))*2 + 1
	if kernelSize < 3 {
		kernelSize = 3
	}
	radius := kernelSize / 2

	kernel := make([][]float64, kernelSize)
	sum := 0.0
	for i := 0; i < kernelSize; i++ {
		kernel[i] = make([]float64, kernelSize)
		for j := 0; j < kernelSize; j++ {
			x := float64(i - radius)
			y := float64(j - radius)
			value := math.Exp(-(x*x + y*y) / (2 * sigma * sigma))
			kernel[i][j] = value
			sum += value
		}
	}

	for i := 0; i < kernelSize; i++ {
		for j := 0; j < kernelSize; j++ {
			kernel[i][j] /= sum
		}
	}

	result := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var r, g, b, a float64

			for ky := 0; ky < kernelSize; ky++ {
				for kx := 0; kx < kernelSize; kx++ {
					px := x + kx - radius
					py := y + ky - radius

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
					cr, cg, cb, ca := c.RGBA()
					weight := kernel[ky][kx]

					r += float64(cr>>8) * weight
					g += float64(cg>>8) * weight
					b += float64(cb>>8) * weight
					a += float64(ca>>8) * weight
				}
			}

			result.SetRGBA(x, y, color.RGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: uint8(a),
			})
		}
	}

	return result
}
