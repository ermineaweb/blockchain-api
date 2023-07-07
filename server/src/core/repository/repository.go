package repository

import "server/src/core/entity"

type Account interface {
	Get(address string) (entity.Account, error)
	GetAll() []entity.Account
}
