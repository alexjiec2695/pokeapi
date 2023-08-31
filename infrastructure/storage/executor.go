package storage

import (
	"context"
	"pokeapi/entities"
)

type Executor interface {
	CreateUsers(ctx context.Context, user entities.Users) error
	GetLogin(ctx context.Context, user entities.Users) (string, error)
	UpdateUsers(ctx context.Context, user entities.Users) error
	GetUsers(ctx context.Context, id string) (entities.Users, error)
}
