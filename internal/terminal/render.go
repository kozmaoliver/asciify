package terminal

import (
	"fmt"
	"asciify/internal/frame"
	"image/color"
	"os"
)

type BackgroundColor string

const (
	BgNone   BackgroundColor = "none"
	BgBlack  BackgroundColor = "black"
	BgWhite  BackgroundColor = "white"
)

// RenderFrame clears the screen and renders a frame to the terminal.
func RenderFrame(f *frame.Frame, bgColor BackgroundColor, useColor bool) {
	// Clear screen: move cursor to home position and clear entire screen
	fmt.Print("\x1b[H\x1b[2J")

	var bgCode string
	switch bgColor {
	case BgBlack:
		bgCode = "\x1b[40m"
	case BgWhite:
		bgCode = "\x1b[47m"
	case BgNone:
		fallthrough
	default:
		bgCode = ""
	}
	if bgCode != "" {
		fmt.Print(bgCode)
	}

	lastColorCode := ""
	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			ch := f.Get(x, y)
			if ch == 0 {
				ch = ' '
			}

			var currentColorCode string
			if useColor && f.Colors != nil {
				c := f.GetColor(x, y)
				if c != nil {
					currentColorCode = colorToANSI(c)
					if currentColorCode != lastColorCode {
						if currentColorCode != "" {
							if bgCode != "" {
								fmt.Print(bgCode + currentColorCode)
							} else {
								fmt.Print(currentColorCode)
							}
							lastColorCode = currentColorCode
						}
					}
				}
			}

			fmt.Fprintf(os.Stdout, "%c", ch)
		}

		if useColor && lastColorCode != "" {
			if bgCode != "" {
				fmt.Print("\x1b[0m" + bgCode)
			} else {
				fmt.Print("\x1b[0m")
			}
			lastColorCode = ""
		}

		if y < f.Height-1 {
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")

	fmt.Print("\x1b[0m")
}

func colorToANSI(c color.Color) string {
	r, g, b, _ := c.RGBA()

	r8 := r >> 8
	g8 := g >> 8
	b8 := b >> 8

	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r8, g8, b8)
}
