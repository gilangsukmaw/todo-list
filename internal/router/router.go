package router

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-v1/cfg/yaml"
	"go-fiber-v1/internal/server"
	"go-fiber-v1/internal/ucase/contract"
	"go-fiber-v1/internal/ucase/ping"
)

func NewRouter(cfg *yaml.Config) *fiber.App {
	router := fiber.New()

	//usecase
	pingUcase := ping.NewPing(cfg)

	router.Get("/ping", handler(pingUcase))

	return router
}

func handler(ucase contract.UseCase, mdws ...Middlewares) fiber.Handler {
	return func(c *fiber.Ctx) error {
		for _, i := range mdws {
			middlewaresFunc := i()
			//if middlewares errors push the response to array
			if middlewaresFunc.Code != 200 {
				response := middlewaresFunc
				return c.Status(response.Code).JSON(response)
			}
		}

		resp := ucase.Serve(c)
		return c.Status(resp.Code).JSON(resp)
	}
}

type Middlewares func() server.Response
