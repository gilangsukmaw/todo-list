package ping

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-v1/cfg/yaml"
	"go-fiber-v1/internal/server"
	"go-fiber-v1/internal/ucase/contract"
)

type ping struct{}

func NewPing() contract.UseCase {
	return &ping{}
}

func (u *ping) Serve(ctx *fiber.Ctx, cfg *yaml.Config) server.Response {
	return server.Response{Code: 200, Data: cfg.App.Name, Message: "ok"}
}
