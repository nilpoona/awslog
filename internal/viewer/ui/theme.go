package ui

import tui "github.com/marcusolsson/tui-go"

const (
	themeErrorLabel = "label.error"
	themeWarnLabel  = "label.warn"
	themeInfoLabel  = "label.info"
	themeDebugLabel = "label.debug"
)

func newThemeErrorLabel() *tui.Theme {
	t := tui.NewTheme()
	t.SetStyle(themeErrorLabel, tui.Style{
		Fg: tui.ColorRed,
	})
	return t
}

func newThemeWarnLabel() *tui.Theme {
	t := tui.NewTheme()
	t.SetStyle(themeWarnLabel, tui.Style{
		Fg: tui.ColorYellow,
	})
	return t
}

func newThemeInfoLabel() *tui.Theme {
	t := tui.NewTheme()
	t.SetStyle(themeInfoLabel, tui.Style{
		Fg: tui.ColorBlue,
	})
	return t
}

func newThemeDebugLabel() *tui.Theme {
	t := tui.NewTheme()
	t.SetStyle(themeDebugLabel, tui.Style{
		Fg: tui.ColorCyan,
	})
	return t
}
