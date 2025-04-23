package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/config"
	"gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/internal/auth"
)

// Run - запускает приложение
func Run() error {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
		return err
	}

	app := fiber.New()

	_ = auth.RegisterRoutes(app, cfg)

	app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"service": cfg.App.Name,
			"version": cfg.App.Version,
		})
	})

	// Start server
	log.Printf("Starting %s on port %d", cfg.App.Name, cfg.App.Port)
	err = app.Listen(fmt.Sprintf(":%d", cfg.App.Port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return err
	}

	return nil
}
