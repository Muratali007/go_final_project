package club

import (
	"fmt"
	"gateway/pkg/club/pb"
	"gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ClubServiceClient
}

func InitServiceClient(c *config.Config) pb.ClubServiceClient {
	cc, err := grpc.Dial(c.ClubSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewClubServiceClient(cc)
}