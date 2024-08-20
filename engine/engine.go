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

type ComponentLayout struct {
	TextView  *tview.TextView
	FormInput *tview.TextArea
}

type Engine struct {
	ApiKey    string
	Context   context.Context
	Component *ComponentLayout
	Model     *genai.GenerativeModel
}

func NewEngine(apiKey string) *Engine {
	engine := &Engine{
		ApiKey:  apiKey,
		Context: context.Background(),
	}

	engine.Component = &ComponentLayout{
		TextView:  engine.TextView(),
		FormInput: engine.FormInput(),
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

	base := layout.BaseLayout(&layout.Base{
		TextView:  e.Component.TextView,
		FormInput: e.Component.FormInput,
	})

	err := tview.NewApplication().SetRoot(base, true).Run()
	if err != nil {
		panic(err.Error())
	}
}
