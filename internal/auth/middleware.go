package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"

	pb "gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/gen/auth"
)

// AuthMiddlewareConfig представляет конфигурацию middleware аутентификации
type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

// InitAuthMiddleware инициализирует middleware аутентификации
func InitAuthMiddleware(svc *ServiceClient) *AuthMiddlewareConfig {
	return &AuthMiddlewareConfig{svc}
}

// AuthRequired middleware проверяет аутентификацию пользователя
func (c *AuthMiddlewareConfig) AuthRequired(ctx fiber.Ctx) error {
	authorization := ctx.Get("Authorization")

	if authorization == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is missing",
		})
	}

	tokenParts := strings.Split(authorization, "Bearer ")
	if len(tokenParts) < 2 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token format",
		})
	}
	token := tokenParts[1]

	res, err := c.svc.Client.ValidateToken(context.Background(), &pb.ValidateTokenRequest{
		AccessToken: token,
	})
	if err != nil || !res.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Поиск совпадения
	fmt.Println("valid", res.Valid, "userId", res.GetUserId())
	ctx.Locals("userId", res.UserId)

	return ctx.Next()
}
