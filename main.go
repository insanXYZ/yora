package main

import (
	"fmt"
	"os"
	"yora/engine"
)

func main() {
	key := os.Getenv("GEMINI_API_KEY")
	if key == "" {
		fmt.Print("GEMINI_API_KEY environment variable not set\n\nYou can set this with:\n# Linux\n	export GEMINI_API_KEY = yourapikey\n# Windows (btw, i dont know what operating system is this)\n	$Env:GEMINI_API_KEY = yourapikey")
		return
	}

	e := engine.NewEngine(key)
	e.Run()
}
