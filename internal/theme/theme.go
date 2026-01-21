package theme

// Theme defines the interface for ASCII character themes.
type Theme interface {
	// Characters returns the ordered list of characters for luminance mapping.
	// Characters should be ordered from darkest to brightest.
	Characters() []rune

	// BrightestChar returns the brightest ASCII character of the theme
	BrightestChar() rune 
	
	// EdgeChars returns a map of edge direction names to characters.
	// This allows themes to customize edge rendering.
	EdgeChars() map[string]rune
}
