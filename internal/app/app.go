package app

import (
	"fmt"
	"gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/internal/fileProcessing"
	"log"

	"github.com/gofiber/fiber/v3"
	"gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/config"
	"gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/internal/auth"
)

// Run - запускает приложение
func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	app := fiber.New()

	_ = auth.RegisterRoutes(app, cfg)
	_ = fileProcessing.RegisterRoutes(app, cfg)

	app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"service": cfg.App.Name,
			"version": cfg.App.Version,
		})
	})

	log.Printf("Starting %s on port %s", cfg.App.Name, cfg.App.Port)
	err = app.Listen(fmt.Sprintf(":%s", cfg.App.Port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Fatal(app.Listen(cfg.App.Port))
}
