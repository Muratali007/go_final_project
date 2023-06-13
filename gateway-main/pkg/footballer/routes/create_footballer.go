package routes

import (
	"context"
	"gateway/pkg/footballer/pb"
	"gateway/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

var logger = logging.GetLogger()

type CreateFootballerRequestBody struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Price   int64  `json:"price"`
	Club    string `json:"club"`
}

func CreateFootballer(ctx *gin.Context, c pb.FootballerServiceClient) {
	f := CreateFootballerRequestBody{}

	if err := ctx.BindJSON(&f); err != nil {
		logger.Info("Bad Request",err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateFootballer(context.Background(), &pb.CreateFootballerRequest{
		Name:    f.Name,
		Surname: f.Surname,
		Price:   f.Price,
		Club:    f.Club,
	})

	if err != nil {
		logger.Info("Bad Gateway",err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
