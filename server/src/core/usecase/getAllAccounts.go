package usecase

import (
	"server/src/core/entity"
	"server/src/core/repository"
)

type getAllAccounts struct {
	repository repository.AccountRepository
}

func NewGetAllAccounts(repository repository.AccountRepository) *getAllAccounts {
	return &getAllAccounts{repository}
}

func (s *getAllAccounts) Exec() []entity.Account {
	return s.repository.GetAll()
}
