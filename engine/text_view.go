package engine

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Engine) TextView() *tview.TextView {
	textView := tview.NewTextView()
	textView.SetDynamicColors(true)
	textView.SetBorder(true)
	textView.SetRegions(true)
	textView.SetWordWrap(true)
	e.SetInputCaptureTextView(textView)
	return textView
}

func (e *Engine) SetInputCaptureTextView(textView *tview.TextView) {
	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlP:
			e.SetFocus(e.Component.FormInput)
		}
		return event
	})
}
