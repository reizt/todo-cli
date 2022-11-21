package cmd

import (
	"fmt"

	"github.com/reizt/todo/src/utils"
)

type cmdPrompt struct{}

func Init() (*cmdPrompt, error) {
	return &cmdPrompt{}, nil
}

func (cmdPrompt) ConfirmClear() (agreed bool) {
	agree := new(string)

	utils.PrintYellow("Are you sure to clear all todos? (y/N): ")
	fmt.Scanf("%s", agree)
	return *agree == "y"
}
