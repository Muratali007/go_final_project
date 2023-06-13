package footballer

import (
	"gateway/pkg/auth"
	"gateway/pkg/config"
	"gateway/pkg/footballer/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/footballer")
	routes.Use(a.AuthRequired)
	routes.POST("/",svc.CreateFootballer)
	routes.GET("/:id", svc.FindOne)
	routes.DELETE("/:id",svc.DeleteOne)
}

func(svc *ServiceClient) CreateFootballer(ctx *gin.Context) {
	routes.CreateFootballer(ctx,svc.Client)
}

func(svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx,svc.Client)
}

func(svc *ServiceClient) DeleteOne(ctx *gin.Context) {
	routes.DeleteOne(ctx,svc.Client)
}
