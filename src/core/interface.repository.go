package core

import (
	"errors"
)

var (
	ErrRepositoryUnexpected      = errors.New("unexpected error")
	ErrRepositoryNotFound        = errors.New("todo not found")
	ErrRepositoryFindMany        = errors.New("failed to get todo list")
	ErrRepositoryInsertFailed    = errors.New("failed to insert todo")
	ErrRepositoryUpdateFailed    = errors.New("failed to update todo")
	ErrRepositoryDeleteFailed    = errors.New("failed to delete todo")
	ErrRepositoryDeleteAllFailed = errors.New("failed to delete all todo")
)

type IRepository interface {
	FindById(id string) (*Todo, error)
	FindMany(input IRepositoryFindManyInput) (*([]Todo), error)
	Insert(input IRepositoryInsertInput) (*Todo, error)
	Update(id string, input IRepositoryUpdateInput) error
	Delete(id string) error
	DeleteAll() error
}

type IRepositoryFindManyInput struct {
	// Add later
}

type IRepositoryInsertInput struct {
	ID          string
	Title       *string
	Description *string
	IsCompleted *bool
}

type IRepositoryUpdateInput struct {
	Title       *string
	Description *string
	IsCompleted *bool
}
