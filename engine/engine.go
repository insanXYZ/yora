package engine

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/rivo/tview"
	"google.golang.org/api/option"
	"yora/layout"
	"yora/throw"
)

type Engine struct {
	Context   context.Context
	Component *layout.ComponentLayout
	App       *tview.Application
	Model     *genai.GenerativeModel
	ApiKey    string
}

func NewEngine(apiKey string) *Engine {
	engine := &Engine{
		App:     tview.NewApplication(),
		ApiKey:  apiKey,
		Context: context.Background(),
	}

	engine.Component = &layout.ComponentLayout{
		FormInput: engine.FormInput(),
		TextView:  engine.TextView(),
		Hint:      engine.Hint(),
	}

	err := engine.InitModel()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return engine
}

func (e *Engine) InitModel() error {
	client, err := genai.NewClient(e.Context, option.WithAPIKey(e.ApiKey))
	if err != nil {
		return throw.ClientGeminiKey()
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	e.Model = model
	return nil
}

func (e *Engine) Run() {

	base := layout.BaseLayout(e.Component)

	err := e.App.SetRoot(base, true).Run()
	if err != nil {
		panic(err.Error())
	}
}

func (e *Engine) QueueUpdateDraw(f func()) {
	go e.App.QueueUpdateDraw(f)
}

func (e *Engine) SetFocus(primitive tview.Primitive) {
	e.QueueUpdateDraw(func() {
		e.App.SetFocus(primitive)
	})
}
