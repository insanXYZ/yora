package engine

import (
	"github.com/rivo/tview"
	"yora/layout"
)

type ComponentLayout struct {
	TextView  *tview.TextView
	FormInput *tview.TextArea
}

type Engine struct {
	ApiKey    string
	Component *tview.Flex
}

func NewEngine(apiKey string) *Engine {
	engine := &Engine{
		ApiKey: apiKey,
	}

	engine.Component = layout.BaseLayout(&layout.Base{
		TextView:  engine.TextView(),
		FormInput: engine.FormInput(),
	})

	return engine
}

func (e *Engine) Run() {

	err := tview.NewApplication().SetRoot(e.Component, true).SetFocus(e.Component).Run()
	if err != nil {
		panic(err.Error())
	}
}
