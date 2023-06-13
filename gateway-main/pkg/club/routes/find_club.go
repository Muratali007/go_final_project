package routes

import (
	"context"
	"gateway/pkg/club/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindClub(ctx *gin.Context, c pb.ClubServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		logger.Info("Bad Gateway",err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.FindClub(context.Background(),&pb.FindClubRequest{
		Id: int64(id),
	})

	if err != nil {
		logger.Info("Bad Gateway",err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
