package engine

import (
	"errors"
	"fmt"
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
	t.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			if t.GetText() != "" {
				iter := e.Model.GenerateContentStream(e.Context, genai.Text(t.GetText()))
				for {
					resp, err := iter.Next()
					if err != nil && errors.Is(err, iterator.Done) {
						break
					}

					s := fmt.Sprint(resp.Candidates[0].Content.Parts[0])

					_, err = e.Component.TextView.Write([]byte(s))
					if err != nil {
						panic(err.Error())
					}

				}
			}
		}
		return event
	})
}
