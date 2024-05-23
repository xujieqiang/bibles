package theme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Mytheme struct {
}

func (t *Mytheme) Font(s fyne.TextStyle) fyne.Resource {

	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
	}
	return nil
}
