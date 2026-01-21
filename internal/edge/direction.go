package edge

import (
	"math"
	"github.com/kozmaoliver/asciify/internal/theme"
)

// EdgeChar maps an edge direction (in radians) to an appropriate ASCII character.
// Direction angles:
// ~0° (horizontal) → '-'
// ~90° (vertical) → '`'
// ~45° → '/'
// ~135° → '\'
func EdgeChar(direction float64) rune {
	// Normalize angle to [0, 2π)
	angle := math.Mod(direction+2*math.Pi, 2*math.Pi)
	
	// Convert to degrees for easier comparison
	deg := angle * 180.0 / math.Pi

	edgeChars := theme.NewDefaultTheme().EdgeChars()
	
	// Map to character based on angle ranges
	// We use ranges around each cardinal direction
	if deg >= 337.5 || deg < 22.5 || (deg >= 157.5 && deg < 202.5) {
		// Horizontal vector => vertical edge
		if edge, ok := edgeChars["vertical"]; ok {
			return edge
		}
		return '-'
	} else if deg >= 67.5 && deg < 112.5 || (deg >= 247.5 && deg < 292.5) {
		// Vertical vactor => horizontal edge
		if edge, ok := edgeChars["horizontal"]; ok {
			return edge
		}
		return '|'
	} else if deg >= 22.5 && deg < 67.5 || (deg >= 202.5 && deg < 247.5) {
		// Diagonal edges (45° and 225°)
		if edge, ok := edgeChars["diagonal1"]; ok {
			return edge
		}
		return '/'
	} else {
		// Diagonal edges (135° and 315°)
		if edge, ok := edgeChars["vdiagonal2"]; ok {
			return edge
		}
		return '\\'
	}
}
