package router

import (
	"chat/api"
	"chat/service"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200,"pong")
		})
		v1.POST("user/register", api.UserRegister)
		v1.GET("ws",service.WsHandler)
	}
	return r
}
