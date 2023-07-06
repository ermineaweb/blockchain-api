package repository

import "server/src/core/entity"

type AccountRepository interface {
	GetOne(address string) (entity.Account, error)
	GetAll() []entity.Account
}
