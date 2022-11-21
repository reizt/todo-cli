package cmd

import (
	"errors"
	"os"

	"github.com/reizt/todo/src/core"
)

var (
	ErrTooFewArguments = errors.New("too few arguments")
)

type Controller struct {
	App core.App
}

func (controller *Controller) Exec(osArgs []string) {
	// recognize command args

	if len(osArgs) == 1 {
		controller.App.Help()
		return
	}

	firstArg := osArgs[1]
	switch firstArg {
	case "-h", "--help":
		controller.App.Help()
	case "-v", "--version":
		controller.App.Version()
	case "list":
		input := newListInput(osArgs)
		controller.App.List(*input)
	case "add":
		input := newAddInput(osArgs)
		controller.App.Add(*input)
	case "mod":
		input := newModInput(osArgs)
		controller.App.Mod(*input)
	case "del":
		input := newDelInput(osArgs)
		controller.App.Del(*input)
	case "clear":
		controller.App.Clear()
	default:
		controller.App.Help()
	}
}

/*
newListInput

	todo list [options]
		options:
			-v, --verbose Detailed list
*/
func newListInput(osArgs []string) *core.ListInput {
	input := core.ListInput{}
	return &input
}

/*
newAddInput

	todo add <title> [options]
	options:
		-d, --description Description
*/
func newAddInput(osArgs []string) *core.AddInput {
	if len(os.Args) < 3 {
		return &core.AddInput{}
	}

	title := (*string)(nil)
	description := (*string)(nil)

	for i, arg := range osArgs {
		if arg == "-t" || arg == "--title" {
			title = &osArgs[i+1]
		}
		if arg == "-d" || arg == "--description" {
			description = &osArgs[i+1]
		}
	}
	input := core.AddInput{
		Title:       title,
		Description: description,
	}
	return &input
}

/*
newModInput

	todo mod <id> [options]
	options:
		-t, --title       Title
		-d, --description Description
*/
func newModInput(osArgs []string) *core.ModInput {
	if len(os.Args) < 3 {
		return &core.ModInput{}
	}

	id := osArgs[2]
	title := (*string)(nil)
	description := (*string)(nil)

	for i, arg := range osArgs {
		if arg == "-t" || arg == "--title" {
			title = &osArgs[i+1]
		}
		if arg == "-d" || arg == "--description" {
			description = &osArgs[i+1]
		}
	}

	input := core.ModInput{
		ID:          id,
		Title:       title,
		Description: description,
	}
	return &input
}

/*
newDelInput

	todo del <id>
*/
func newDelInput(osArgs []string) *core.DelInput {
	if len(os.Args) < 3 {
		return &core.DelInput{}
	}

	id := osArgs[2]
	input := core.DelInput{
		ID: id,
	}
	return &input
}
