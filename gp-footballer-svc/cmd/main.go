package main

import (
	"fmt"
	"football/pkg/config"
	"football/pkg/db"
	"football/pkg/logging"
	"football/pkg/pb"
	"football/pkg/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

var logger = logging.GetLogger()

func main()  {
	conn, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed at config:",err)
	}

	h := db.Init(conn.DBUrl)

	lis, err := net.Listen("tcp", conn.Port)
	if err != nil {
		log.Fatalln("Failed to listening", err)
	}

	fmt.Println("Listennig in this port:", conn.Port)

	s := services.Server{
		H: h,
	}

	grpcServer :=grpc.NewServer()
	pb.RegisterFootballerServiceServer(grpcServer,&s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
