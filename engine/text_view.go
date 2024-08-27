package engine

import (
	"errors"
	"fmt"
	"yora/color"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"google.golang.org/api/iterator"
)

func (e *Engine) TextView() *tview.TextView {

	textView := tview.NewTextView()
	textView.SetDynamicColors(true)
	textView.SetBorder(true)
	textView.SetRegions(true)
	textView.SetWordWrap(true)
	textView.SetBorderColor(color.RED)
	textView.SetBackgroundColor(color.BLACK)
	e.SetInputCaptureTextView(textView)
	e.SetHub("textview")

	go e.ListenTextviewHub(textView)

	return textView
}

func (e *Engine) ListenTextviewHub(comp *tview.TextView) {

	for {
		select {
		case data := <-e.Hub["textview"]:
			if data.Status == SENDMESSAGE {

				e.QueueUpdateDraw(func() {
					comp.Write([]byte(fmt.Sprint("👤 [green]You:\n[white]" + data.Data + "\n\n")))
					comp.Write([]byte("🤖 [blue]Yora:\n[white]"))
				})

				go func() {
					iter := e.Model.QuestionStream(data.Data)
					for {
						resp, err := iter.Next()
						if err != nil && errors.Is(err, iterator.Done) {
							break
						}

						s := fmt.Sprint(resp.Candidates[0].Content.Parts[0])
						e.QueueUpdateDraw(func() {
							comp.Write([]byte(s))
							comp.ScrollToEnd()
						})
					}
					comp.Write([]byte("\n"))

					e.SendToHub("forminput", Hub{
						Status: SETSTATUSSENDER,
					})

				}()

			}
		}
	}
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
