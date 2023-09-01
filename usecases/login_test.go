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

func TestLogin_Login_Successful(t *testing.T) {
	storage := new(mock.ExecutorStorageMock)
	ctx := context.Background()
	expect := "123456789"
	users := entities.Users{
		ID:        "",
		Name:      "",
		Email:     "",
		Password:  "",
		Address:   "",
		Birthdate: "",
		City:      "",
	}

	uc := usecases.NewLogin(storage)

	storage.On("GetLogin", ctx, users).Return(expect, nil)

	id, err := uc.Login(ctx, users)

	assert.NoError(t, err)
	assert.NotNil(t, id)
	assert.Equal(t, expect, id)
}

func TestLogin_Login_Successful_With_Errors(t *testing.T) {
	storage := new(mock.ExecutorStorageMock)
	ctx := context.Background()
	expect := "123456789"
	users := entities.Users{
		ID:        "",
		Name:      "",
		Email:     "",
		Password:  "",
		Address:   "",
		Birthdate: "",
		City:      "",
	}

	uc := usecases.NewLogin(storage)

	storage.On("GetLogin", ctx, users).Return(expect, errors.New("error"))

	_, err := uc.Login(ctx, users)

	assert.Error(t, err)

}
