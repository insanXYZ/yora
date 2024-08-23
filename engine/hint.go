package engine

import "github.com/rivo/tview"

func (e *Engine) Hint() *tview.TextView {
	text := tview.NewTextView()
	text.SetDynamicColors(true)
	text.SetText("[green]ctrl + p = [white]change tab | [green]ctrl + space = [white]enter message")
	return text
}
