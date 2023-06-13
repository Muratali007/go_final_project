package routes

import (
	"context"
	"gateway/pkg/footballer/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteOne(ctx *gin.Context, c pb.FootballerServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		logger.Info("Bad Gateway",err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	_, err = c.DeleteFootballer(context.Background(),&pb.DeleteFootballerRequest{
		Id: int64(id),
	})

	if err != nil {
		logger.Info("Bad Gateway",err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
