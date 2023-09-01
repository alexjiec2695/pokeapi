package di

import (
	"pokeapi/infrastructure/handlers"
	"pokeapi/infrastructure/rest"
	resty2 "pokeapi/infrastructure/resty"
	"pokeapi/infrastructure/storage"
	"pokeapi/usecases"
)

func Start() error {
	server := rest.NewServer()
	str := ""

	persistence, err := storage.NewConnectionStorage(str)
	if err != nil {
		return err
	}
	createUsers := usecases.NewCreateUsers(persistence)
	login := usecases.NewLogin(persistence)
	updateUsers := usecases.NewUpdateUsers(persistence)
	getUsers := usecases.NewGetUsers(persistence)
	resty := resty2.NewPokemon()
	pokes := usecases.NewPokemon(resty)

	routers := handlers.NewHandler(server, createUsers, login, updateUsers, getUsers, pokes)

	return routers.Start()
}
