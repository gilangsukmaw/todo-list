package contract

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-v1/internal/server"
)

type UseCase interface {
	Serve(r *fiber.Ctx) server.Response
}
