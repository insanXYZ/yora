package engine

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (e *Engine) FormInput() *tview.TextArea {
	textarea := tview.NewTextArea()
	textarea.SetBorder(true)
	textarea.SetTitle("Message...")
	textarea.SetTitleAlign(tview.AlignLeft)
	e.setInputCaptureFormInput(textarea)

	e.SetHub("forminput")
	return textarea
}

func (e *Engine) setInputCaptureFormInput(comp *tview.TextArea) {
	sending := false

	comp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlSpace:
			message := comp.GetText()

			if message != "" && !sending {

				if err := e.CheckConnection(); err != nil {
					comp.SetText("no internet connection", false)
					return event
				}

				e.SendToHub("textview", Hub{
					Data:   message,
					Status: SENDMESSAGE,
				})

				sending = true

				go func() {
					for {
						select {
						case <-e.Hub["forminput"]:
							e.QueueUpdateDraw(func() {
								comp.SetText("", false)
							})
							sending = false
							return
						default:
							e.QueueUpdateDraw(func() {
								comp.SetText("wait..", false)
							})
						}
					}
				}()

			}
		case tcell.KeyCtrlP:
			e.SetFocus(e.Component.TextView)
			comp.SetText("", false)
		}
		return event
	})
}
