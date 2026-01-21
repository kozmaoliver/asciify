package converter

import (
	"github.com/kozmaoliver/asciify/internal/edge"
	"github.com/kozmaoliver/asciify/internal/theme"
)

// Resolver decides which character to use for each pixel
// based on luminance and edge information.
type Resolver struct {
	Theme      theme.Theme
	EdgeCutoff float64
}

func NewResolver(t theme.Theme, edgeCutoff float64) *Resolver {
	return &Resolver{
		Theme:      t,
		EdgeCutoff: edgeCutoff,
	}
}

func (r *Resolver) Resolve(lum float64, e edge.Edge) rune {
	if e.Strength > r.EdgeCutoff {
		return edge.EdgeChar(e.Direction)
	}

	chars := r.Theme.Characters()
	if len(chars) == 0 {
		return ' '
	}

	index := int(lum * float64(len(chars)))
	if index >= len(chars) {
		index = len(chars) - 1
	}
	if index < 0 {
		index = 0
	}

	return chars[index]
}
