package edge

import (
	"math"
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
	
	// Map to character based on angle ranges
	// We use ranges around each cardinal direction
	if deg >= 337.5 || deg < 22.5 || (deg >= 157.5 && deg < 202.5) {
		// Horizontal edges (0° and 180°)
		return '-'
	} else if deg >= 67.5 && deg < 112.5 || (deg >= 247.5 && deg < 292.5) {
		// Vertical edges (90° and 270°)
		return '`'
	} else if deg >= 22.5 && deg < 67.5 || (deg >= 202.5 && deg < 247.5) {
		// Diagonal edges (45° and 225°)
		return '/'
	} else {
		// Diagonal edges (135° and 315°)
		return '\\'
	}
}
