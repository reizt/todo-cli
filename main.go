package main

import (
	"fmt"
	"os"

	controller "github.com/reizt/todo/src/adapters/controller/cmd"
	prompt "github.com/reizt/todo/src/adapters/prompt/cmd"
	"github.com/reizt/todo/src/adapters/renderer/console"
	"github.com/reizt/todo/src/adapters/repository/sqlite"
	"github.com/reizt/todo/src/core"
	"github.com/reizt/todo/src/utils"
)

func main() {
	repository, err := sqlite.Init()
	if err != nil {
		utils.PrintRed(fmt.Sprintf("ERROR: unexpected error occured!\n%s\n", err.Error()))
		return
	}
	renderer, err := console.Init()
	if err != nil {
		utils.PrintRed(fmt.Sprintf("ERROR: unexpected error occured!\n%s\n", err.Error()))
		return
	}
	prompt, err := prompt.Init()
	if err != nil {
		utils.PrintRed(fmt.Sprintf("ERROR: unexpected error occured!\n%s\n", err.Error()))
		return
	}

	app := core.App{
		Repository: repository,
		Renderer:   renderer,
		Prompt:     prompt,
	}
	controller := controller.Controller{App: app}

	controller.Exec(os.Args)

	defer repository.Close()
}
