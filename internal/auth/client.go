package auth

import (
	"fmt"

	"google.golang.org/grpc"

	"gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/config"
	pb "gitlab.crja72.ru/golang/2025/spring/course/projects/go21/api-gateway/gen/auth"
)

type ServiceClient struct {
	Client pb.AuthClient
}

func InitServiceClient(c *config.Config) pb.AuthClient {
	cc, err := grpc.Dial(c.Microservices.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}

	return pb.NewAuthClient(cc)
}
