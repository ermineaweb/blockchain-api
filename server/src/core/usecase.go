package core

type Usecase interface {
	Exec(any) any
}
