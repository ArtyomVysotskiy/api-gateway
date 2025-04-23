package main

import (
	"log"

	"gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
