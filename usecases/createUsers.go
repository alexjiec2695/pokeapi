package usecases

import (
	"context"
	"github.com/google/uuid"
	"pokeapi/entities"
	"pokeapi/infrastructure/storage"
)

type CreateUsers struct {
	storage storage.Executor
}

type CreateUsersExecutor interface {
	CreateUsers(ctx context.Context, users entities.Users) error
}

func NewCreateUsers(storage storage.Executor) CreateUsersExecutor {
	return &CreateUsers{
		storage: storage,
	}
}

func (c *CreateUsers) CreateUsers(ctx context.Context, users entities.Users) error {
	users.ID = uuid.New().String()
	return c.storage.CreateUsers(ctx, users)
}
