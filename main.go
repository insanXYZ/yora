package main

import (
	"fmt"
	"os"
	"yora/engine"
	"yora/throw"
)

func main() {

	key := os.Getenv("GEMINI_API_KEY")
	if key == "" {
		fmt.Println(throw.MissingKey())
		return
	}
	e := engine.NewEngine(key)
	e.Run()

}
