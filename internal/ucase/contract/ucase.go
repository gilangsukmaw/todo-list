package contract

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/todo-list-app1/todo-list-backend/cfg/yaml"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
)

type UseCase interface {
	Serve(r *fiber.Ctx, cfg *yaml.Config) server.Response
}
