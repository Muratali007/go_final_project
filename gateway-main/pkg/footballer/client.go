package footballer

import (
	"gateway/pkg/config"
	"gateway/pkg/footballer/pb"
	"gateway/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var logger = logging.GetLogger()

type ServiceClient struct {
	Client pb.FootballerServiceClient
}

func InitServiceClient(c *config.Config) pb.FootballerServiceClient {
	cc, err := grpc.Dial(c.FootballerSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logger.Info("Could not connect:", err)
	}

	return pb.NewFootballerServiceClient(cc)
}
