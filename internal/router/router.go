package router

import (
	"github.com/gofiber/fiber/v2"
	dbConn "gitlab.com/todo-list-app1/todo-list-backend/cfg/db"
	"gitlab.com/todo-list-app1/todo-list-backend/cfg/yaml"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/middlewares"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/repositories"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/contract"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/ping"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/todos"
)

func NewRouter(cfg *yaml.Config) *fiber.App {
	router := fiber.New()

	db := dbConn.NewDatabase(cfg)

	//repositories
	var (
		//userRepo = repositories.NewUser(db.Db)
		todoRepo = repositories.NewTodo(db.Db)
	)

	//middlewares
	var (
		tesMdwr  = middlewares.TesMdwr
		tesMdwr2 = middlewares.TesMdwr2
	)

	//ucase
	var (
		pingUcase       = ping.NewPing()
		AllTodosUcase   = todos.NewAllTodos(todoRepo)
		CreateTodoUcase = todos.NewCreateTodo(todoRepo)
		DoneTodoUcase   = todos.NewDoneTodos(todoRepo)
		UndoneTodoUcase = todos.NewUndoneTodos(todoRepo)
		DeleteTodoUcase = todos.NewDeleteTodos(todoRepo)
	)

	//group
	api := router.Group("/api")
	v1 := api.Group("/v1")

	router.Get("/ping/:param", handler(cfg, pingUcase, tesMdwr, tesMdwr2))

	//v1.Get("/users", handler(cfg, usersUcase))

	//todo crud
	v1.Get("/todos", handler(cfg, AllTodosUcase))
	v1.Post("/todos", handler(cfg, CreateTodoUcase))
	v1.Put("/todos/done/:id", handler(cfg, DoneTodoUcase))
	v1.Put("/todos/undone/:id", handler(cfg, UndoneTodoUcase))
	v1.Delete("/todos/delete/:id", handler(cfg, DeleteTodoUcase))

	return router
}

func handler(cfg *yaml.Config, ucase contract.UseCase, mdws ...Middlewares) fiber.Handler {
	return func(c *fiber.Ctx) error {
		for _, i := range mdws {
			middlewaresFunc := i(c, cfg)
			if middlewaresFunc.Code != 200 {
				response := middlewaresFunc
				return c.Status(response.Code).JSON(response)
			}
		}

		resp := ucase.Serve(c, cfg)
		return c.Status(resp.Code).JSON(resp)
	}
}

type Middlewares func(c *fiber.Ctx, cfg *yaml.Config) server.Response
