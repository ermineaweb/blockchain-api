package inmemory

import (
	"errors"
	"server/src/core/entity"
)

type InMemoryRepository struct{}

var accounts = []entity.Account{{Address: "abc", Balance: 15}, {Address: "def", Balance: 20}}

func (r InMemoryRepository) Get(address string) (entity.Account, error) {
	for _, account := range accounts {
		if account.Address == address {
			return account, nil
		}
	}
	return entity.Account{}, errors.New("account not found")
}

func (r InMemoryRepository) GetAll() []entity.Account {
	return accounts
}
