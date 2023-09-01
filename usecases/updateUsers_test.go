package usecases_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"pokeapi/entities"
	"pokeapi/usecases"
	"pokeapi/usecases/mock"
	"testing"
)

func TestUpdateUsers_UpdateUsers_Success(t *testing.T) {
	storage := new(mock.ExecutorStorageMock)
	ctx := context.Background()
	user := entities.Users{}

	storage.On("UpdateUsers", ctx, user).Return(nil)
	update := usecases.NewUpdateUsers(storage)

	err := update.UpdateUsers(ctx, user)

	assert.NoError(t, err)
}

func TestUpdateUsers_UpdateUsers_Success_With_Error(t *testing.T) {
	storage := new(mock.ExecutorStorageMock)
	ctx := context.Background()
	user := entities.Users{}

	storage.On("UpdateUsers", ctx, user).Return(errors.New("error"))
	update := usecases.NewUpdateUsers(storage)

	err := update.UpdateUsers(ctx, user)

	assert.Error(t, err)
}
