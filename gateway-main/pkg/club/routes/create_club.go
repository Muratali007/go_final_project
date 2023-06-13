package routes

import (
	"context"
	"gateway/pkg/club/pb"
	"gateway/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)
var logger = logging.GetLogger()

type CreateClubRequestBody struct {
	Name string `json:"name"`
	Year int64 `json:"year"`
	Title int64 `json:"title"`
}

func CreateClub(ctx *gin.Context, c pb.ClubServiceClient) {
	b := CreateClubRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		logger.Info("Bad Request",err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateClub(context.Background(), &pb.CreateClubRequest{
		Name: b.Name,
		Year: b.Year,
		Trophy: b.Title,
	})

	if err != nil {
		logger.Info("Bad Gateway",err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
