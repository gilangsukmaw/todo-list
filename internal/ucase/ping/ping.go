package ping

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/todo-list-app1/todo-list-backend/cfg/yaml"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/contract"
)

type ping struct{}

func NewPing() contract.UseCase {
	return &ping{}
}

func (u *ping) Serve(ctx *fiber.Ctx, cfg *yaml.Config) server.Response {
	return server.Response{Code: 200, Data: cfg.App.Name, Message: "ok"}
}
