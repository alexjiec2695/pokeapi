package usecases

import (
	"context"
	"pokeapi/entities"
	"pokeapi/infrastructure/storage"
)

type UpdateUsers struct {
	storage storage.Executor
}

type UpdateUsersExecutor interface {
	UpdateUsers(ctx context.Context, users entities.Users) error
}

func NewUpdateUsers(storage storage.Executor) UpdateUsersExecutor {
	return &UpdateUsers{
		storage: storage,
	}
}

func (c *UpdateUsers) UpdateUsers(ctx context.Context, users entities.Users) error {
	return c.storage.UpdateUsers(ctx, users)
}
