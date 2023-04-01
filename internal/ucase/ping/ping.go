package ping

import (
	"github.com/gofiber/fiber/v2"
	cfg "gitlab.com/todo-list-app1/todo-list-backend/cfg/env"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/contract"
)

type ping struct{}

func NewPing() contract.UseCase {
	return &ping{}
}

func (u *ping) Serve(ctx *fiber.Ctx, cfg *cfg.Config) server.Response {
	return server.Response{Code: 200, Data: cfg.Name, Message: "ok"}
}
