package main

import (
	"os"

	controller "github.com/reizt/todo/src/adapters/controller/cmd"
	prompt "github.com/reizt/todo/src/adapters/prompt/cmd"
	"github.com/reizt/todo/src/adapters/renderer/console"
	"github.com/reizt/todo/src/adapters/repository/sqlite"
	"github.com/reizt/todo/src/core"
)

func main() {
	repository, _ := sqlite.Init()
	renderer, _ := console.Init()
	prompt, _ := prompt.Init()

	app := core.App{
		Repository: repository,
		Renderer:   renderer,
		Prompt:     prompt,
	}
	controller := controller.Controller{App: app}

	controller.Exec(os.Args)

	defer repository.Close()
}
