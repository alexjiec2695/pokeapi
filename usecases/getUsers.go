package usecases

import (
	"context"
	"pokeapi/entities"
	"pokeapi/infrastructure/storage"
)

type GetUsers struct {
	storage storage.Executor
}

type GetUsersExecutor interface {
	GetUsers(ctx context.Context, id string) (entities.Users, error)
}

func NewGetUsers(storage storage.Executor) GetUsersExecutor {
	return &GetUsers{
		storage: storage,
	}
}

func (c *GetUsers) GetUsers(ctx context.Context, id string) (entities.Users, error) {
	return c.storage.GetUsers(ctx, id)
}
