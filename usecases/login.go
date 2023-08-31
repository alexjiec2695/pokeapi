package usecases

import (
	"context"
	"pokeapi/entities"
	"pokeapi/infrastructure/storage"
)

type Login struct {
	storage storage.Executor
}

type LoginExecutor interface {
	Login(ctx context.Context, users entities.Users) (string, error)
}

func NewLogin(storage storage.Executor) LoginExecutor {
	return &Login{
		storage: storage,
	}
}

func (c *Login) Login(ctx context.Context, users entities.Users) (string, error) {
	return c.storage.GetLogin(ctx, users)
}
