package usecases_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pokeapi/entities"
	"pokeapi/usecases"
	ucMock "pokeapi/usecases/mock"
	"testing"
)

func TestCreateUsers_CreateUsers_Successful(t *testing.T) {
	store := new(ucMock.ExecutorStorageMock)
	ctx := context.Background()
	user := entities.Users{}

	store.On("CreateUsers", ctx, mock.Anything).Return(nil)

	uc := usecases.NewCreateUsers(store)

	err := uc.CreateUsers(ctx, user)

	assert.NoError(t, err)
}
