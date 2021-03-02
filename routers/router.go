package routers

import (
	"github.com/gin-gonic/gin"

	"gin-gonic/middleware/jwt"
	"gin-gonic/pkg/setting"
	v1 "gin-gonic/routers/api/v1"
)

// InitRouter initialize the server with with selected configuration
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", v1.Ping)
		apiv1.POST("/login", v1.Login)
	}

	apiv1.Use(jwt.Check())
	{
		apiv1.GET("/secure/ping", v1.SecurePing)
	}

	return r
}
