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

type GetFilesResponse struct {
	FileId    string
	FileName  string
	FileSize  string
	MimeType  string
	Extension string
	CreateAt  string
}

type SearchFilesRequest struct {
	FileId     string `json:"file_id"`
	SearchTerm string `json:"search_term"`
	UserId     string `json:"user_id"`
}

type GetFilesByIdRequest struct {
	IdUser string `json:"id_user"`
	FileId string `json:"file_id"`
}

func UploadFile(c fiber.Ctx, client *ServiceClientFileProcessing) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Файл не получен")
	}

	fmt.Printf("file %+v\n", fileHeader)

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Не удалось открыть файл")
	}
	defer file.Close()
	fmt.Printf("open %+v\n", file)

	fileSize := fileHeader.Size
	if fileSize == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Файл пустой")
	}

	// Подключаемся к gRPC
	stream, err := client.Client.UploadFile(c.Context())
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка подключения к gRPC")
	}

	buf := make([]byte, 1024*16) // 16KB чанки, чтобы снизить нагрузку
	first := true
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Printf("Error reading file: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Ошибка при чтении файла")
		}

		if n == 0 {
			// Если файл закончился, завершаем передачу
			break
		}

		fmt.Printf("Uploading chunk: %d bytes\n", n)

		req := &pb.UploadFileRequest{
			Chunk: buf[:n],
		}

		if first {
			req.Filename = fileHeader.Filename
			first = false
		}

		if err := stream.Send(req); err != nil {
			fmt.Printf("Error sending chunk: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Ошибка при отправке чанка")
		}
	}

	// Завершаем и получаем ответ
	resp, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("Error closing stream: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка завершения gRPC стрима")
	}

	fmt.Printf("File upload completed: %s\n", resp.Message)
	return c.SendString(fmt.Sprintf("Файл %s загружен. Ответ: %s", fileHeader.Filename, resp.Message))
}

func GetFiles(c fiber.Ctx, client *ServiceClientFileProcessing) error {
	res, err := client.Client.GetFiles(c.Context(), &pb.GetFilesRequest{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("GetFiles")
	}
	req := GetFilesResponse{
		FileId:    res.FileId,
		FileName:  res.FileName,
		FileSize:  res.FileSize,
		MimeType:  res.MimeType,
		Extension: res.Extension,
		CreateAt:  res.CreateAt,
	}

	return c.SendString(fmt.Sprintf("%+v", req))
}

func GetFilesById(c fiber.Ctx, client *ServiceClientFileProcessing) error {
	var req GetFilesByIdRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	res, err := client.Client.GetFileByID(c.Context(), &pb.GetFileByIDRequest{
		UserId: req.IdUser,
		FileId: req.FileId,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("GetFiles")
	}
	return c.SendString(fmt.Sprintf("%+v", res))
}

func SearchFiles(c fiber.Ctx, client *ServiceClientFileProcessing) error {
	fmt.Println("SearchFiles")
	var req SearchFilesRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	res, err := client.Client.SearchFile(c.Context(), &pb.SearchFileRequest{
		FileId:     req.FileId,
		SearchTerm: req.SearchTerm,
		UserId:     req.UserId,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("GetFiles")
	}

	return c.SendString(fmt.Sprintf("%+v", res))
}

func RegisterRoutes(app *fiber.App, c *config.Config) *ServiceClientFileProcessing {
	svc := &ServiceClientFileProcessing{
		Client: InitServiceClient(c),
	}

	fileProcessing := app.Group("/fileProcessing")

	fileProcessing.Post("/UploadFile", func(c fiber.Ctx) error {
		return UploadFile(c, svc)
	})
	fileProcessing.Get("/GetFiles", func(c fiber.Ctx) error {
		return GetFiles(c, svc)
	})
	fileProcessing.Post("/GetFilesById", func(c fiber.Ctx) error {
		return GetFilesById(c, svc)
	})
	fileProcessing.Post("/SearchFiles", func(c fiber.Ctx) error {
		return SearchFiles(c, svc)
	})

	return svc
}
