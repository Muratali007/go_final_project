package routes

import (
	"context"
	"gateway/pkg/auth/pb"
	"gateway/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

var logger = logging.GetLogger()

type LoginRequestBody struct {
	Name string 	`json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		logger.Info("Bad Request",err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		logger.Info("Bad Gateway",err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

