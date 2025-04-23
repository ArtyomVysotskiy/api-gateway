package auth

import (
	"github.com/gofiber/fiber/v3"
	"gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/config"
	pb "gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/gen/auth"
)

// Request types
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type ValidateTokenRequest struct {
	AccessToken string `json:"access_token"`
}

type LogoutRequest struct {
	AccessToken string `json:"access_token"`
}

// Handlers
func Register(c fiber.Ctx, client *ServiceClient) error {
	var req RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := client.Client.Register(c.Context(), &pb.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id": res.UserId,
	})
}

func Login(c fiber.Ctx, client *ServiceClient) error {
	var req LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := client.Client.Login(c.Context(), &pb.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  res.AccessToken,
		"refresh_token": res.RefreshToken,
	})
}

func RefreshToken(c fiber.Ctx, client *ServiceClient) error {
	var req RefreshTokenRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := client.Client.RefreshToken(c.Context(), &pb.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": res.AccessToken,
	})
}

func ValidateToken(c fiber.Ctx, client *ServiceClient) error {
	var req ValidateTokenRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := client.Client.ValidateToken(c.Context(), &pb.ValidateTokenRequest{
		AccessToken: req.AccessToken,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"valid": res.Valid,
	})
}

func Logout(c fiber.Ctx, client *ServiceClient) error {
	var req LogoutRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err := client.Client.Logout(c.Context(), &pb.LogoutRequest{
		AccessToken: req.AccessToken,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}

// Route registration
func RegisterRoutes(app *fiber.App, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	// Создаем группу маршрутов для аутентификации
	authGroup := app.Group("/auth")

	// Регистрируем все обработчики
	authGroup.Post("/register", func(c fiber.Ctx) error {
		return Register(c, svc)
	})

	authGroup.Post("/login", func(c fiber.Ctx) error {
		return Login(c, svc)
	})

	authGroup.Post("/refresh", func(c fiber.Ctx) error {
		return RefreshToken(c, svc)
	})

	authGroup.Post("/validate", func(c fiber.Ctx) error {
		return ValidateToken(c, svc)
	})

	authGroup.Post("/logout", func(c fiber.Ctx) error {
		return Logout(c, svc)
	})

	return svc
}
