package core

import (
	"github.com/reizt/todo/src/utils"
)

const (
	cliVersion = "v1.0.0"
)

type App struct {
	Repository IRepository
	Renderer   IRenderer
	Prompt     IPrompt
}

func (app *App) validateTodo(todo Todo) (IsValid bool) {
	if err := todo.IsValid(); err != nil {
		if err == ErrTitleIsRequired {
			app.Renderer.TitleIsRequired()
			return
		} else if err == ErrTitleIsTooLong {
			app.Renderer.TitleIsTooLong(titleMaxLength)
			return
		}
		app.Renderer.InvalidInput()
		return
	}
	return true
}

func (app *App) Version() {
	app.Renderer.Version(cliVersion)
}

func (app *App) Help() {
	app.Renderer.Help()
}

type ListInput struct {
	OnlyUnFinished bool
}

func (app *App) List(input ListInput) {
	repositoryFindManyInput := IRepositoryFindManyInput{}
	if input.OnlyUnFinished {
		repositoryFindManyInput.Finished = utils.Ptr(false)
	}
	todos, err := app.Repository.FindMany(repositoryFindManyInput)
	if err != nil {
		app.Renderer.ListFailedToFindMany(err)
		return
	}

	app.Renderer.List(*todos)
}

type AddInput struct {
	Title       *string
	Description *string
	IsFinished  *bool
}

func (app *App) Add(input AddInput) {
	id := newId()
	IsFinished := false
	if input.IsFinished != nil && *input.IsFinished {
		IsFinished = true
	}
	todo := Todo{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
		IsFinished:  &IsFinished,
	}
	isValid := app.validateTodo(todo)
	if !isValid {
		app.Renderer.AddHelp()
		return
	}

	addedTodo, err := app.Repository.Insert(IRepositoryInsertInput(todo))
	if err != nil {
		app.Renderer.AddFailedToInsert(err)
		return
	}
	app.Renderer.AddSucceeded(*addedTodo)

	todos, err := app.Repository.FindMany(IRepositoryFindManyInput{})
	if err != nil {
		app.Renderer.ListFailedToFindMany(err)
		return
	}
	app.Renderer.List(*todos)
}

type ModInput struct {
	ID          string
	Title       *string
	Description *string
	IsFinished  *bool
}

func (app *App) Mod(input ModInput) {
	if input.ID == "" {
		app.Renderer.ModHelp()
		return
	}

	todo, err := app.Repository.FindById(input.ID)
	if err != nil {
		if err == ErrRepositoryNotFound {
			app.Renderer.NotFound(input.ID)
			return
		}
		app.Renderer.UnexpectedError()
		return
	}

	if input.Title != nil {
		todo.Title = input.Title
	}
	if input.Description != nil {
		todo.Description = input.Description
	}
	if input.IsFinished != nil {
		todo.IsFinished = input.IsFinished
	}
	isValid := app.validateTodo(*todo)
	if !isValid {
		app.Renderer.ModHelp()
		return
	}

	err = app.Repository.Update(todo.ID, IRepositoryUpdateInput{
		Title:       input.Title,
		Description: input.Description,
		IsFinished:  input.IsFinished,
	})
	if err != nil {
		app.Renderer.ModFailedToUpdate(err)
		return
	}
	app.Renderer.ModSucceeded(*todo)

	todos, err := app.Repository.FindMany(IRepositoryFindManyInput{})
	if err != nil {
		app.Renderer.ListFailedToFindMany(err)
		return
	}
	app.Renderer.List(*todos)
}

type DelInput struct {
	ID string
}

func (app *App) Del(input DelInput) {
	if input.ID == "" {
		app.Renderer.DelHelp()
		return
	}

	todo, err := app.Repository.FindById(input.ID)
	if err != nil {
		if err == ErrRepositoryNotFound {
			app.Renderer.NotFound(input.ID)
			return
		}
		app.Renderer.UnexpectedError()
		return
	}

	err = app.Repository.Delete(todo.ID)
	if err != nil {
		app.Renderer.DelFailed(err)
		return
	}

	app.Renderer.DelSucceeded(*todo)

	todos, err := app.Repository.FindMany(IRepositoryFindManyInput{})
	if err != nil {
		app.Renderer.ListFailedToFindMany(err)
		return
	}
	app.Renderer.List(*todos)
}

func (app *App) Clear() {
	agreed := app.Prompt.ConfirmClear()
	if !agreed {
		app.Renderer.ClearCanceled()
		return
	}

	err := app.Repository.DeleteAll()
	if err != nil {
		app.Renderer.ClearFailed(err)
		return
	}

	app.Renderer.ClearSucceeded()
}

type FinInput struct {
	ID string
}

func (app *App) Fin(input FinInput) {
	if input.ID == "" {
		app.Renderer.FinHelp()
		return
	}

	todo, err := app.Repository.FindById(input.ID)
	if err != nil {
		if err == ErrRepositoryNotFound {
			app.Renderer.NotFound(input.ID)
			return
		}
		app.Renderer.UnexpectedError()
		return
	}

	err = app.Repository.Update(todo.ID, IRepositoryUpdateInput{
		IsFinished: utils.Ptr(true),
	})
	if err != nil {
		app.Renderer.FinFailed(err)
		return
	}

	app.Renderer.FinSucceeded(*todo)

	todos, err := app.Repository.FindMany(IRepositoryFindManyInput{})
	if err != nil {
		app.Renderer.ListFailedToFindMany(err)
		return
	}
	app.Renderer.List(*todos)
}
