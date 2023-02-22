package ping

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-v1/cfg/yaml"
	"go-fiber-v1/internal/server"
	"go-fiber-v1/internal/ucase/contract"
)

type ping struct {
	cfg *yaml.Config
}

func NewPing(cfg *yaml.Config) contract.UseCase {
	return &ping{
		cfg: cfg,
	}
}

func (u *ping) Serve(ctx *fiber.Ctx) server.Response {
	return server.Response{Code: 200, Data: u.cfg.App.Name, Message: "ok"}
}
