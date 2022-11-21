package core

import (
	"errors"
)

type Todo struct {
	ID          string
	Title       *string
	Description *string
}

const (
	titleMaxLength = 180
)

var (
	ErrTitleIsRequired = errors.New("title is required")
	ErrTitleIsTooLong  = errors.New("title is too long")
)

func (todo *Todo) IsValid() error {
	if todo.Title == nil || *todo.Title == "" {
		return ErrTitleIsRequired
	}
	if len(*todo.Title) > titleMaxLength {
		return ErrTitleIsTooLong
	}
	return nil
}
