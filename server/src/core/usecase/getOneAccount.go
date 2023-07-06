package usecase

import (
	"server/src/core/entity"
	"server/src/core/repository"
)

type getOneAccount struct {
	repository repository.AccountRepository
}

func NewGetOneAccount(repository repository.AccountRepository) *getOneAccount {
	return &getOneAccount{repository}
}

func (s *getOneAccount) Exec(address string) (entity.Account, error) {
	return s.repository.GetOne(address)
}
