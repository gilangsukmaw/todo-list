package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/todo-list-app1/todo-list-backend/cfg/yaml"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
)

func TesMdwr(ctx *fiber.Ctx, cfg *yaml.Config) server.Response {
	fmt.Println("ctx --> ", ctx.AllParams())
	return server.Response{Code: 400, Message: "test 1"}
}

func TesMdwr2(ctx *fiber.Ctx, cfg *yaml.Config) server.Response {
	return server.Response{Code: 200, Message: "test 2"}
}
