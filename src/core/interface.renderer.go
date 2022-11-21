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
	ClearSucceeded()
	ClearFailed(err error)
	ClearCanceled()
	DelFailed(err error)
}
