package fileProcessing

import (
	"fmt"
	"gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/config"
	pb "gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/gen/fileProcessing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClientFileProcessing struct {
	Client pb.FileProcessingClient
}

func InitServiceClient(c *config.Config) pb.FileProcessingClient {
	cc, err := grpc.NewClient(c.Microservices.FileProcessingSvcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}
	return pb.NewFileProcessingClient(cc)
}
