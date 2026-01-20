package luminance

import (
	"image/color"
)

// Luminance calculates the perceived brightness of a color.
// Uses the formula: L = 0.2126*R + 0.7152*G + 0.0722*B
func Luminance(c color.Color) float64 {
	r, g, b, _ := c.RGBA()
	
	// Convert from 16-bit to 8-bit and normalize to 0.0-1.0
	rf := float64(r>>8) / 255.0
	gf := float64(g>>8) / 255.0
	bf := float64(b>>8) / 255.0

	lum := 0.2126*rf + 0.7152*gf + 0.0722*bf
	
	if lum > 1.0 {
		lum = 1.0
	}
	if lum < 0.0 {
		lum = 0.0
	}
	
	return lum
}
