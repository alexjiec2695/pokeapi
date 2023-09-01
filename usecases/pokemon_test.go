package usecases_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"pokeapi/entities"
	"pokeapi/usecases"
	"pokeapi/usecases/mock"
	"testing"
)

func TestPokemon_GetPokemon_Success(t *testing.T) {
	restyMock := new(mock.ExecutorRestyPokemon)
	ctx := context.Background()
	expect := entities.PokemonDetail{}

	restyMock.On("GetPokemonDetail", "url").Return(expect, nil)

	uc := usecases.NewPokemon(restyMock)

	pokemons, err := uc.GetPokemon(ctx, "url")

	assert.NoError(t, err)
	assert.NotNil(t, pokemons)
	assert.Equal(t, expect, pokemons)
}
