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

func TestGetUsers_GetUsers_Successful(t *testing.T) {
	storage := new(mock.ExecutorStorageMock)
	ctx := context.Background()
	id := "1234566789"
	expect := entities.Users{
		ID:        "1231445",
		Name:      "Alexander",
		Email:     "A@gmauil.com",
		Password:  "XXXXXXX",
		Address:   "carrera ",
		Birthdate: "26-07-1995",
		City:      "Medellin",
	}

	storage.On("GetUsers", ctx, id).Return(expect, nil)

	uc := usecases.NewGetUsers(storage)

	user, err := uc.GetUsers(ctx, id)

	assert.NoError(t, err)
	assert.Equal(t, expect, user)

}

func TestGetUsers_GetUsers_Successful_With_Error(t *testing.T) {
	storage := new(mock.ExecutorStorageMock)
	ctx := context.Background()
	id := "1234566789"
	expect := entities.Users{
		ID:        "1231445",
		Name:      "Alexander",
		Email:     "A@gmauil.com",
		Password:  "XXXXXXX",
		Address:   "carrera ",
		Birthdate: "26-07-1995",
		City:      "Medellin",
	}
	storage.On("GetUsers", ctx, id).Return(expect, errors.New("error"))
	uc := usecases.NewGetUsers(storage)
	_, err := uc.GetUsers(ctx, id)
	assert.Error(t, err)
}
