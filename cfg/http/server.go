package http

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-v1/cfg/yaml"
	"log"
	_ "net/http"
	_ "net/http/pprof"
)

func Run(cfg *yaml.Config) {
	//start server

	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", hello)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
