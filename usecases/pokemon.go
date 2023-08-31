package usecases

import (
	"context"
	"pokeapi/entities"
	"pokeapi/infrastructure/resty"
)

type Pokemon struct {
	resty resty.ExecutorPokemon
}

type PokemonExecutor interface {
	GetPokemons(ctx context.Context) ([]entities.PokemonItem, error)
	GetPokemon(ctx context.Context, url string) (entities.PokemonDetail, error)
}

func NewPokemon(resty resty.ExecutorPokemon) PokemonExecutor {
	return &Pokemon{
		resty: resty,
	}
}

func (c *Pokemon) GetPokemons(ctx context.Context) ([]entities.PokemonItem, error) {
	p := entities.PokemonList{}
	channel := make(chan []entities.PokemonItem)

	go c.resty.GetPokemon(channel)

	for v := range channel {
		p.Results = append(p.Results, v...)
	}

	return p.Results, nil
}

func (c *Pokemon) GetPokemon(ctx context.Context, url string) (entities.PokemonDetail, error) {
	return c.resty.GetPokemonDetail(url)
}
