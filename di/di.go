package di

import (
	"pokeapi/infrastructure/handlers"
	"pokeapi/infrastructure/rest"
	"pokeapi/infrastructure/storage"
	"pokeapi/usecases"
)

func Start() error {
	server := rest.NewServer()
	str := "tpr77y14irv178s44t8u:pscale_pw_xiKe5hNoHKpGGtBJxY5i8q4adQmPMjr8GYPt9Uq8MdK@tcp(aws.connect.psdb.cloud)/test?tls=true&interpolateParams=true"

	persistence, err := storage.NewConnectionStorage(str)
	if err != nil {
		return err
	}
	createUsers := usecases.NewCreateUsers(persistence)
	login := usecases.NewLogin(persistence)
	updateUsers := usecases.NewUpdateUsers(persistence)
	getUsers := usecases.NewGetUsers(persistence)

	routers := handlers.NewRouters(server, createUsers, login, updateUsers, getUsers)

	return routers.Start()
}
