package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-fiber-v1/cfg/yaml"
	"go-fiber-v1/internal/router"
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
