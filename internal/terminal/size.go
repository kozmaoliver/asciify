package terminal

import (
	"os"
	"golang.org/x/sys/unix"
)

// Size represents terminal dimensions in characters.
type Size struct {
	Width  int
	Height int
}

func GetTerminalSize() (Size, error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return Size{Width: 80, Height: 24}, err
	}

	return Size{
		Width:  int(ws.Col),
		Height: int(ws.Row),
	}, nil
}
