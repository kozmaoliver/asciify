package theme

// DefaultTheme implements the default 10-level ASCII theme.
type DefaultTheme struct{}

func NewDefaultTheme() *DefaultTheme {
	return &DefaultTheme{}
}

func (t *DefaultTheme) Characters() []rune {
	return []rune{' ', '.', ':', '-', '=', '+', '*', '%', '@', '#'}
}

func (t *DefaultTheme) BrightestChar() rune {
	chars := t.Characters()
	return chars[len(chars)-1]
}

// TODO: make it a struct
func (t *DefaultTheme) EdgeChars() map[string]rune {
	return map[string]rune{
		"horizontal": '-',
		"vertical":   '|',
		"diagonal1":  '/',
		"diagonal2":  '\\',
	}
}
