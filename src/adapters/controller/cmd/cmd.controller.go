package cmd

import (
	"errors"
	"os"

	"github.com/reizt/todo/src/utils"

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

	switch osArgs[1] {
	case "-h", "--help":
		controller.App.Help()
	case "-v", "--version":
		controller.App.Version()
	case "list", "l":
		input := newListInput(osArgs)
		controller.App.List(*input)
	case "add", "a":
		input := newAddInput(osArgs)
		controller.App.Add(*input)
	case "mod", "m":
		input := newModInput(osArgs)
		controller.App.Mod(*input)
	case "fin", "f":
		input := newFinInput(osArgs)
		controller.App.Fin(*input)
	case "del", "d":
		input := newDelInput(osArgs)
		controller.App.Del(*input)
	case "clear", "c":
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
	OnlyUnFinished := false

	for i := 0; i < len(osArgs); i++ {
		if osArgs[i] == "-u" {
			OnlyUnFinished = true
		}
	}
	input := core.ListInput{
		OnlyUnFinished: OnlyUnFinished,
	}
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

	Title := (*string)(nil)
	Description := (*string)(nil)
	IsFinished := (*bool)(nil)

	for i := 0; i < len(osArgs); i++ {
		if (osArgs[i] == "-t" || osArgs[i] == "--title") && len(osArgs) >= i+2 {
			Title = &osArgs[i+1]
			i++
			continue
		}
		if (osArgs[i] == "-d" || osArgs[i] == "--description") && len(osArgs) >= i+2 {
			Description = &osArgs[i+1]
			i++
			continue
		}
		if osArgs[i] == "-c" || osArgs[i] == "--completed" {
			if len(osArgs) >= i+2 && osArgs[i+1] == "false" {
				IsFinished = utils.Ptr(false)
				i++
			} else {
				IsFinished = utils.Ptr(true)
			}
			continue
		}
	}
	input := core.AddInput{
		Title:       Title,
		Description: Description,
		IsFinished:  IsFinished,
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

	ID := osArgs[2]
	Title := (*string)(nil)
	Description := (*string)(nil)
	IsFinished := (*bool)(nil)

	for i := 0; i < len(osArgs); i++ {
		if (osArgs[i] == "-t" || osArgs[i] == "--title") && len(osArgs) >= i+2 {
			Title = &osArgs[i+1]
			i++
			continue
		}
		if (osArgs[i] == "-d" || osArgs[i] == "--description") && len(osArgs) >= i+2 {
			Description = &osArgs[i+1]
			i++
			continue
		}
		if osArgs[i] == "-c" || osArgs[i] == "--completed" {
			if len(osArgs) >= i+2 && osArgs[i+1] == "false" {
				IsFinished = utils.Ptr(false)
				i++
			} else {
				IsFinished = utils.Ptr(true)
			}
			continue
		}
	}
	input := core.ModInput{
		ID:          ID,
		Title:       Title,
		Description: Description,
		IsFinished:  IsFinished,
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

/*
newFinInput

	todo fin <id>
*/
func newFinInput(osArgs []string) *core.FinInput {
	if len(os.Args) < 3 {
		return &core.FinInput{}
	}

	id := osArgs[2]
	input := core.FinInput{
		ID: id,
	}
	return &input
}
