package core

import (
	"errors"
)

var (
	ErrRendererUnexpected = errors.New("unexpected error")
)

type IRenderer interface {
	Help()
	AddHelp()
	ModHelp()
	DelHelp()
	FinHelp()
	Version(version string)
	TitleIsRequired()
	TitleIsTooLong(maxLength int)
	InvalidInput()
	UnexpectedError()
	List(todos []Todo)
	ListFailedToFindMany(err error)
	AddSucceeded(todo Todo)
	AddFailedToInsert(err error)
	ModSucceeded(todo Todo)
	ModFailedToUpdate(err error)
	NotFound(id string)
	DelSucceeded(todo Todo)
	DelFailed(err error)
	ClearSucceeded()
	ClearFailed(err error)
	ClearCanceled()
	FinSucceeded(todo Todo)
	FinFailed(err error)
}
