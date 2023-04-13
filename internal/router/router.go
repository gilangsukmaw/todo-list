package router

import (
	"github.com/gofiber/fiber/v2"
	dbConn "gitlab.com/todo-list-app1/todo-list-backend/cfg/db"
	cfg "gitlab.com/todo-list-app1/todo-list-backend/cfg/env"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/middlewares"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/repositories"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/auth"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/contract"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/ping"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/todo_group"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/todos"
)

func NewRouter(cfg *cfg.Config) *fiber.App {
	router := fiber.New()

	db := dbConn.NewDatabase(cfg)

	//repositories
	var (
		userRepo      = repositories.NewUser(db.Db)
		todoRepo      = repositories.NewTodo(db.Db)
		todoGroupRepo = repositories.NewTodoGroup(db.Db)
	)

	//middlewares
	var (
		tesMdwr     = middlewares.TesMdwr
		tesMdwr2    = middlewares.TesMdwr2
		BearerToken = middlewares.ValidateJWT
	)

	//ucase
	var (
		pingUcase       = ping.NewPing()
		AllTodosUcase   = todos.NewAllTodos(todoRepo, todoGroupRepo)
		CreateTodoUcase = todos.NewCreateTodo(todoRepo, todoGroupRepo)
		DoneTodoUcase   = todos.NewDoneTodos(todoRepo)
		UndoneTodoUcase = todos.NewUndoneTodos(todoRepo)
		DeleteTodoUcase = todos.NewDeleteTodos(todoRepo)

		todoGroupsUcase = todo_group.NewAllTodoGroup(todoRepo, todoGroupRepo)
		loginUcase      = auth.NewLogin(userRepo)
	)

	//group
	api := router.Group("/api")
	v1 := api.Group("/v1")

	router.Get("/ping/:param", handler(cfg, pingUcase, tesMdwr, tesMdwr2))

	//v1.Get("/users", handler(cfg, usersUcase))

	//todo group
	v1.Get("/todo-groups", handler(cfg, todoGroupsUcase, BearerToken))

	//todo crud
	v1.Get("/todos/:unique", handler(cfg, AllTodosUcase, BearerToken))
	v1.Post("/todos/:unique", handler(cfg, CreateTodoUcase, BearerToken))
	v1.Put("/todos/done/:id", handler(cfg, DoneTodoUcase, BearerToken))
	v1.Put("/todos/undone/:id", handler(cfg, UndoneTodoUcase, BearerToken))
	v1.Delete("/todos/delete/:id", handler(cfg, DeleteTodoUcase, BearerToken))

	//auth
	v1.Post("/auth/login", handler(cfg, loginUcase))

	return router
}

func handler(cfg *cfg.Config, ucase contract.UseCase, mdws ...Middlewares) fiber.Handler {
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

type Middlewares func(c *fiber.Ctx, cfg *cfg.Config) server.Response
