package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"pokeapi/entities"
	"pokeapi/infrastructure/rest"
	"pokeapi/usecases"
)

type Routers struct {
	app         *fiber.App
	createUsers usecases.CreateUsersExecutor
	updateUsers usecases.UpdateUsersExecutor
	login       usecases.LoginExecutor
	getUsers    usecases.GetUsersExecutor
}

func NewRouters(app *rest.Server,
	createUsers usecases.CreateUsersExecutor,
	login usecases.LoginExecutor,
	updateUsers usecases.UpdateUsersExecutor,
	getUsers usecases.GetUsersExecutor,
) Routers {
	return Routers{
		app:         app.App,
		createUsers: createUsers,
		updateUsers: updateUsers,
		login:       login,
		getUsers:    getUsers,
	}
}

func (r *Routers) CreateUsers() {
	r.app.Post("Users", func(ctx *fiber.Ctx) error {
		users := entities.Users{}

		err := json.Unmarshal(ctx.Body(), &users)
		if err != nil {
			return err
		}

		return r.createUsers.CreateUsers(ctx.Context(), users)
	})
}

func (r *Routers) Login() {
	r.app.Post("Login", func(ctx *fiber.Ctx) error {
		users := entities.Users{}

		err := json.Unmarshal(ctx.Body(), &users)
		if err != nil {
			return err
		}

		id, err := r.login.Login(ctx.Context(), users)
		if err != nil {
			return ctx.SendStatus(http.StatusNoContent)
		}
		type response struct {
			ID string
		}

		b, _ := json.Marshal(response{
			ID: id,
		})

		return ctx.Status(http.StatusOK).Send(b)
	})
}

func (r *Routers) UpdateUsers() {
	r.app.Put("Users", func(ctx *fiber.Ctx) error {
		users := entities.Users{}

		err := json.Unmarshal(ctx.Body(), &users)
		if err != nil {
			return err
		}

		return r.updateUsers.UpdateUsers(ctx.Context(), users)
	})
}

func (r *Routers) GetUsers() {
	r.app.Get("User/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		user, err := r.getUsers.GetUsers(ctx.Context(), id)
		if err != nil {
			return err
		}
		b, err := json.Marshal(user)
		if err != nil {
			return err
		}
		return ctx.Status(200).Send(b)

	})
}

func (r *Routers) Start() error {
	r.CreateUsers()
	r.UpdateUsers()
	r.GetUsers()
	r.Login()
	return r.app.Listen(":3000")
}
