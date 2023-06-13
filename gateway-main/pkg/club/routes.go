package club

import (
	"gateway/pkg/auth"
	"gateway/pkg/club/routes"
	"gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/club")
	routes.Use(a.AuthRequired)
	routes.POST("/",svc.CreateClub)
	routes.GET("/:id",svc.GetClub)
}

func (svc *ServiceClient) CreateClub(ctx *gin.Context) {
	routes.CreateClub(ctx,svc.Client)
}

func (svc *ServiceClient) GetClub(ctx *gin.Context) {
	routes.FindClub(ctx, svc.Client)
}