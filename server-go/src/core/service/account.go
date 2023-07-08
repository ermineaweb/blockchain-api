package service

import (
	"server/src/core/entity"
	"server/src/core/repository"
)

type Account struct {
	repository repository.Account
}

func NewAccountService(repository repository.Account) *Account {
	return &Account{repository}
}

func (s *Account) GetAll() []entity.Account {
	return s.repository.GetAll()
}

func (s *Account) Get(address string) (entity.Account, error) {
	return s.repository.Get(address)
}
