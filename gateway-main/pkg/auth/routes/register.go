package routes

import (
	"context"
	"gateway/pkg/auth/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	b := RegisterRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		logger.Info("Bad Request",err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    b.Email,
		Password: b.Password,
		Name: b.Name,
	})

	if err != nil {
		logger.Info("Bad Gateway",err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
