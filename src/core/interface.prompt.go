package core

type IPrompt interface {
	ConfirmClear() (agreed bool)
}
