package core

type ICommand interface {
	Execute(params []string)
}
