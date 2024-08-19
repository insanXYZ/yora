package engine

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Engine) FormInput() *tview.TextArea {
	textarea := tview.NewTextArea()
	textarea.SetTitle("Message...")
	textarea.SetBorder(true)
	textarea.SetTitleAlign(tview.AlignLeft)
	return textarea
}

func (e *Engine) setInputCaptureFormInput(t *tview.TextArea) {
	t.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			if t.GetTextLength() != 0 {
				
			}
		}
		return event
	})
}
