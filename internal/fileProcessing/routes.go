package fileProcessing

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/config"
	pb "gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/gen/fileProcessing"
	"io"
)

type UploadFileRequest struct {
	Filename string
	Chunk    []byte
}

func UploadFile(c fiber.Ctx, client *ServiceClientFileProcessing) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Файл не получен")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Не удалось открыть файл")
	}
	defer file.Close()

	// Подключаемся к gRPC
	stream, err := client.Client.UploadFile(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка подключения к gRPC")
	}

	buf := make([]byte, 1024*32) // 32KB чанки
	first := true
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Ошибка при чтении файла")
		}

		req := &pb.UploadFileRequest{
			Chunk: buf[:n],
		}

		// 🔸 Добавляем имя только в первом чанке
		if first {
			req.Filename = fileHeader.Filename
			first = false
		}

		if err := stream.Send(req); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Ошибка при отправке чанка")
		}
	}

	// Завершаем и получаем ответ
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка завершения gRPC стрима")
	}

	return c.SendString(fmt.Sprintf("Файл %s загружен. Ответ: %s", fileHeader.Filename, resp.Message))
}

func RegisterRoutes(app *fiber.App, c *config.Config) *ServiceClientFileProcessing {
	svc := &ServiceClientFileProcessing{
		Client: InitServiceClient(c),
	}

	fileProcessing := app.Group("/fileProcessing")

	fileProcessing.Post("/UploadFile", func(c fiber.Ctx) error {
		return UploadFile(c, svc)
	})

	return svc
}
