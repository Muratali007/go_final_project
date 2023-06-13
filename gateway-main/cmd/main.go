package main

import (
	"gateway/pkg/auth"
	"gateway/pkg/club"
	"gateway/pkg/config"
	"gateway/pkg/footballer"
	"gateway/pkg/logging"
	"github.com/gin-gonic/gin"
)

var logger = logging.GetLogger()

func main() {
	conn, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed at config",err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r,&conn)
	footballer.RegisterRoutes(r,&conn,&authSvc)
	club.RegisterRoutes(r,&conn,&authSvc)

	r.Run(conn.Port)
}
