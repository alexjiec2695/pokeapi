package resty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pokeapi/entities"
)

type pokemon struct {
}

type ExecutorPokemon interface {
	GetPokemon(c chan<- []entities.PokemonItem)
	GetPokemonDetail(url string) (entities.PokemonDetail, error)
}

func NewPokemon() ExecutorPokemon {
	return &pokemon{}
}

func (p *pokemon) GetPokemon(c chan<- []entities.PokemonItem) {
	defer func() {
		close(c)
	}()
	offset := 0
	limit := 100
	var apiResult entities.PokemonList
	for {
		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?offset=%v&limit=%v", offset, limit)
		result, err := http.Get(url)
		if err != nil {
			continue
		}

		if err := json.NewDecoder(result.Body).Decode(&apiResult); err != nil {
			continue
		}

		if len(apiResult.Results) == 0 {
			break
		}

		c <- apiResult.Results

		offset += limit
	}
}

func (p *pokemon) GetPokemonDetail(url string) (entities.PokemonDetail, error) {
	detail := entities.PokemonDetail{}
	result, err := http.Get(url)
	if err != nil {
		return entities.PokemonDetail{}, nil
	}

	if err := json.NewDecoder(result.Body).Decode(&detail); err != nil {
		return entities.PokemonDetail{}, err
	}

	return detail, nil
}
