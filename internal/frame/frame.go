package frame

import (
	"image/color"
)

// Frame represents a rendered frame of ASCII characters.
type Frame struct {
	Width  int
	Height int
	Cells  [][]rune
	Colors [][]color.Color
}

func New(width, height int) *Frame {
	cells := make([][]rune, height)
	for i := range cells {
		cells[i] = make([]rune, width)
	}
	return &Frame{
		Width:  width,
		Height: height,
		Cells:  cells,
		Colors: nil,
	}
}

func (f *Frame) EnableColors() {
	if f.Colors == nil {
		f.Colors = make([][]color.Color, f.Height)
		for i := range f.Colors {
			f.Colors[i] = make([]color.Color, f.Width)
		}
	}
}

func (f *Frame) SetColor(x, y int, c color.Color) {
	if f.Colors != nil && x >= 0 && x < f.Width && y >= 0 && y < f.Height {
		f.Colors[y][x] = c
	}
}

func (f *Frame) GetColor(x, y int) color.Color {
	if f.Colors != nil && x >= 0 && x < f.Width && y >= 0 && y < f.Height {
		return f.Colors[y][x]
	}
	return nil
}

func (f *Frame) Set(x, y int, ch rune) {
	if x >= 0 && x < f.Width && y >= 0 && y < f.Height {
		f.Cells[y][x] = ch
	}
}

func (f *Frame) Get(x, y int) rune {
	if x >= 0 && x < f.Width && y >= 0 && y < f.Height {
		return f.Cells[y][x]
	}
	return ' '
}
