package routes

import (
	"context"
	"gateway/pkg/footballer/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindOne(ctx *gin.Context, c pb.FootballerServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		logger.Info("Bad Gateway",err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		logger.Info("Bad Gateway",err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
