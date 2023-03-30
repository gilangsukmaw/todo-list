package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gitlab.com/todo-list-app1/todo-list-backend/cfg/yaml"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/router"
	"log"
	_ "net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
)

func Run(cfg *yaml.Config) {
	//start server

	// Fiber instance
	app := fiber.New()

	// Default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	//setup router
	app.Mount("/", router.NewRouter(cfg))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	// ...

	if err := app.Listen(":8081"); err != nil {
		log.Panic(err)
	}
}
