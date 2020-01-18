package ui

import (
	"fmt"

	tui "github.com/marcusolsson/tui-go"
)

type (
	themeName string
)

const (
	themeDefault themeName = "default"
	themeError   themeName = "error"
	themeWarn    themeName = "warn"
	themeInfo    themeName = "info"
	themeDebug   themeName = "debug"
)

func labelStyleName(name themeName) string {
	return fmt.Sprintf("label.%s", name.String())
}

func (t themeName) String() string {
	return string(t)
}

func newDefaultStyle() tui.Style {
	return tui.Style{
		Fg: tui.ColorWhite,
		Bg: tui.ColorBlack,
	}
}

func newErrorStyle() tui.Style {
	return tui.Style{
		Fg: tui.ColorRed,
		Bg: tui.ColorDefault,
	}
}

func newWarnStyle() tui.Style {
	return tui.Style{
		Fg: tui.ColorYellow,
		Bg: tui.ColorDefault,
	}
}

func newInfoStyle() tui.Style {
	return tui.Style{
		Fg: tui.ColorBlue,
		Bg: tui.ColorDefault,
	}
}

func newDebugStyle() tui.Style {
	return tui.Style{
		Fg: tui.ColorCyan,
	}
}

func newTheme() *tui.Theme {
	t := tui.NewTheme()
	t.SetStyle(labelStyleName(themeError), newErrorStyle())
	t.SetStyle(labelStyleName(themeInfo), newInfoStyle())
	t.SetStyle(labelStyleName(themeDebug), newDebugStyle())
	t.SetStyle(labelStyleName(themeWarn), newWarnStyle())
	t.SetStyle(labelStyleName(themeDefault), newDebugStyle())
	t.SetStyle("list.item", newDefaultStyle())
	t.SetStyle("list.item.selected", tui.Style{
		Fg: tui.ColorBlack,
		Bg: tui.ColorWhite,
	})
	t.SetStyle("normal", newDefaultStyle())
	return t
}
