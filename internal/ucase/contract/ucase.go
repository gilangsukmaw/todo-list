package contract

import (
	"github.com/gofiber/fiber/v2"
	cfg "gitlab.com/todo-list-app1/todo-list-backend/cfg/env"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
)

type UseCase interface {
	Serve(r *fiber.Ctx, cfg *cfg.Config) server.Response
}
