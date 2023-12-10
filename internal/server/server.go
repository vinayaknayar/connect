package server

import (
	"flag"
	"os"
	"time"

	"connect/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

var (
	addr = flag.String("addr", ":"+os.Getenv("PORT"), "http service address")
	cert = flag.String("cert", "", "path to TLS certificate")
	key  = flag.String("key", "", "path to TLS key")
)

func Run() error {
	flag.Parse()

	if *addr == ":" {
		*addr = ":8080"
	}

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", )
	app.Get("/room/:id", )
	app.Get("/room/:id/ws", )
	app.Get("/room/:id/chat", )
	app.Get("/room/:id/chat/ws", )
	app.Get("/room/:id/viewer/ws", )
	app.Get("/stream/:sid", )
	app.Get("/stream/:sid/ws", )
	app.Get("/stream/:sid/chat/ws", )
	app.Get("/stream/:sid/viewer/ws", )

	return app.Listen(*addr)
}