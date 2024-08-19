package engine

import "github.com/rivo/tview"

func (e *Engine) TextView() *tview.TextView {
	textView := tview.NewTextView()
	textView.SetDynamicColors(true)
	textView.SetBorder(true)
	return textView
}
