package engine

import (
	"fmt"
	"yora/color"

	"github.com/rivo/tview"
)

func (e *Engine) Hint() *tview.TextView {
	text := tview.NewTextView()
	text.SetDynamicColors(true)
	text.SetBackgroundColor(color.BLACK)
	text.SetTextColor(color.WHITE)
	text.SetText(fmt.Sprintf("[green]ctrl + p = [%s]change tab | [green]ctrl + space = [%s]enter message", color.WHITE, color.WHITE))
	return text
}
