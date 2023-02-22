package contract

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-v1/cfg/yaml"
	"go-fiber-v1/internal/server"
)

type UseCase interface {
	Serve(r *fiber.Ctx, cfg *yaml.Config) server.Response
}
