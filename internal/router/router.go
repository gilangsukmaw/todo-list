package router

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-v1/cfg/yaml"
	"go-fiber-v1/internal/middlewares"
	"go-fiber-v1/internal/server"
	"go-fiber-v1/internal/ucase/contract"
	"go-fiber-v1/internal/ucase/ping"
)

func NewRouter(cfg *yaml.Config) *fiber.App {
	router := fiber.New()

	//db := dbConn.NewDatabase(cfg)

	//repositories

	//middlewares
	var (
		tesMdwr  = middlewares.TesMdwr
		tesMdwr2 = middlewares.TesMdwr2
	)

	//ucase
	var (
		pingUcase = ping.NewPing()
	)

	router.Get("/ping/:param", handler(cfg, pingUcase, tesMdwr, tesMdwr2))

	return router
}

func handler(cfg *yaml.Config, ucase contract.UseCase, mdws ...Middlewares) fiber.Handler {
	return func(c *fiber.Ctx) error {
		for _, i := range mdws {
			middlewaresFunc := i(c, cfg)
			//if middlewares errors push the response to array
			if middlewaresFunc.Code != 200 {
				response := middlewaresFunc
				return c.Status(response.Code).JSON(response)
			}
		}

		resp := ucase.Serve(c, cfg)
		return c.Status(resp.Code).JSON(resp)
	}
}

type Middlewares func(c *fiber.Ctx, cfg *yaml.Config) server.Response
