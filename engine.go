package main

import "github.com/rivo/tview"

type Engine struct {
	ApiKey string
	Tview  *tview.Application
}

func NewEngine(apiKey string) *Engine {
	return &Engine{ApiKey: apiKey}
}

func (e *Engine) Run() {

}
