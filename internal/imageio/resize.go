package imageio

import (
	"image"
	"math"
)

// Terminal characters are taller than wide.
const CharAspectRatio = 0.5

// ResizeForTerminal resizes an image to fit within terminal bounds
func ResizeForTerminal(img image.Image, termWidth, termHeight int) image.Image {
	bounds := img.Bounds()
	imgWidth := bounds.Dx()
	imgHeight := bounds.Dy()

	imgAspect := float64(imgWidth) / float64(imgHeight)

	effectiveTermHeight := float64(termHeight) / CharAspectRatio
	termAspect := float64(termWidth) / effectiveTermHeight

	var newWidth, newHeight int
	if imgAspect > termAspect {
		newWidth = termWidth
		newHeight = int(float64(termWidth) / imgAspect * CharAspectRatio)
	} else {
		newHeight = termHeight
		newWidth = int(float64(termHeight) * imgAspect / CharAspectRatio)
	}

	if newWidth > termWidth {
		newWidth = termWidth
		newHeight = int(float64(termWidth) / imgAspect * CharAspectRatio)
	}
	if newHeight > termHeight {
		newHeight = termHeight
		newWidth = int(float64(termHeight) * imgAspect / CharAspectRatio)
	}

	resized := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	scaleX := float64(imgWidth) / float64(newWidth)
	scaleY := float64(imgHeight) / float64(newHeight)

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			srcX := int(math.Floor(float64(x) * scaleX))
			srcY := int(math.Floor(float64(y) * scaleY))
			
			if srcX >= imgWidth {
				srcX = imgWidth - 1
			}
			if srcY >= imgHeight {
				srcY = imgHeight - 1
			}

			resized.Set(x, y, img.At(bounds.Min.X+srcX, bounds.Min.Y+srcY))
		}
	}

	return resized
}
