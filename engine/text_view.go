package engine

import "github.com/rivo/tview"

func (e *Engine) TextView() *tview.TextView {
	textView := tview.NewTextView()
	textView.SetDynamicColors(true)
	textView.SetBorder(true)
	textView.SetRegions(true)
	textView.SetWordWrap(true)
	return textView
}
