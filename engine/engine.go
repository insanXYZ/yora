package engine

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"yora/layout"
	thirdparty "yora/third_party"

	"github.com/rivo/tview"
)

type Engine struct {
	Context   context.Context
	Component *layout.ComponentLayout
	App       *tview.Application
	Model     *thirdparty.GenaiAI
	Hub       map[string]chan Hub
}

func NewEngine(apiKey string) *Engine {
	ctx := context.Background()

	engine := &Engine{
		App: tview.NewApplication(),
		Hub: make(map[string]chan Hub),
	}

	engine.Component = &layout.ComponentLayout{
		FormInput: engine.FormInput(),
		TextView:  engine.TextView(),
		Hint:      engine.Hint(),
	}

	model, err := thirdparty.NewGenai(ctx, apiKey)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	engine.Model = model

	return engine
}

func (e *Engine) Run() {

	base := layout.BaseLayout(e.Component)

	err := e.App.SetRoot(base, true).Run()
	if err != nil {
		panic(err.Error())
	}
}

func (e *Engine) SetHub(key string) {
	e.Hub[key] = make(chan Hub)
}

func (e *Engine) SendToHub(key string, payload Hub) {
	e.Hub[key] <- payload
}

func (e *Engine) CheckConnection() (err error) {
	_, err = http.Get("https://8.8.8.8")
	return
}

func (e *Engine) QueueUpdateDraw(f func()) {
	e.App.QueueUpdateDraw(f)
	time.Sleep(100 * time.Millisecond)
}

func (e *Engine) SetFocus(primitive tview.Primitive) {
	go e.QueueUpdateDraw(func() {
		e.App.SetFocus(primitive)
	})
}
