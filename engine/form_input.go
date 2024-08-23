package engine

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/google/generative-ai-go/genai"
	"github.com/rivo/tview"
	"google.golang.org/api/iterator"
)

func (e *Engine) FormInput() *tview.TextArea {
	textarea := tview.NewTextArea()
	textarea.SetBorder(true)
	textarea.SetTitle("Message...")
	textarea.SetTitleAlign(tview.AlignLeft)
	e.setInputCaptureFormInput(textarea)
	return textarea
}

func (e *Engine) setInputCaptureFormInput(t *tview.TextArea) {
	status := false

	t.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlSpace:

			if t.GetText() != "" && !status {

				if _, err := http.Get("https://8.8.8.8"); err != nil {
					t.SetText("no internet connection", false)
					return event
				}

				message := fmt.Sprint("👤 [green]You:\n[white]" + t.GetText() + "\n\n")

				e.Component.TextView.Write([]byte(message))

				status = true

				text := t.GetText()
				ch := make(chan bool)

				go func() {
					iter := e.Model.GenerateContentStream(e.Context, genai.Text(text))
					e.Component.TextView.Write([]byte("🤖 [blue]Yora:\n[white]"))
					for {
						resp, err := iter.Next()
						if err != nil && errors.Is(err, iterator.Done) {
							break
						}

						s := fmt.Sprint(resp.Candidates[0].Content.Parts[0])

						e.App.QueueUpdateDraw(func() {
							e.Component.TextView.Write([]byte(s))
							e.Component.TextView.ScrollToEnd()
						})

						time.Sleep(100 * time.Millisecond)
					}
					e.Component.TextView.Write([]byte("\n"))
					ch <- true
				}()

				go func() {
					for {
						select {
						case <-ch:
							close(ch)
							e.App.QueueUpdateDraw(func() {
								t.SetText("", false)
							})
							time.Sleep(100 * time.Millisecond)
							status = false
							return
						default:
							e.App.QueueUpdateDraw(func() {
								t.SetText("wait..", false)
							})
							time.Sleep(100 * time.Millisecond)
						}
					}
				}()

			}
		case tcell.KeyCtrlP:
			e.SetFocus(e.Component.TextView)
			t.SetText("", false)
		}
		return event
	})
}
