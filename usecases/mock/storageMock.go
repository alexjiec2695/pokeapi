package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"pokeapi/entities"
)

type ExecutorStorageMock struct {
	mock.Mock
}

func (m *ExecutorStorageMock) CreateUsers(ctx context.Context, user entities.Users) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *ExecutorStorageMock) GetLogin(ctx context.Context, user entities.Users) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}

func (m *ExecutorStorageMock) UpdateUsers(ctx context.Context, user entities.Users) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *ExecutorStorageMock) GetUsers(ctx context.Context, id string) (entities.Users, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entities.Users), args.Error(1)
}

type ExecutorRestyPokemon struct {
	mock.Mock
}

func (m *ExecutorRestyPokemon) GetPokemon(c chan<- []entities.PokemonItem) {
	m.Called(c)
	return
}

func (m *ExecutorRestyPokemon) GetPokemonDetail(url string) (entities.PokemonDetail, error) {
	args := m.Called(url)
	return args.Get(0).(entities.PokemonDetail), args.Error(1)
}
