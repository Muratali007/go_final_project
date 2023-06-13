package main

import (
	"auth/pkg/config"
	"auth/pkg/db"
	"auth/pkg/logging"
	"auth/pkg/pb"
	"auth/pkg/services"
	"auth/pkg/utils"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var logger = logging.GetLogger()

func main() {
	conn, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed at config",err)
	}

	h := db.Init(conn.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       conn.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", conn.Port)
	if err != nil {
		log.Fatalln("Failed to listening", err)
	}

	fmt.Println("Auth Svc on", conn.Port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer,&s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("failed to serve", err)
	}
}
