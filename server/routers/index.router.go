package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/controllers"
)

func SetUpIndexRouter(router *gin.RouterGroup) {

	indexController := controllers.IndexController{}

	index := router.Group("/")
	{
		index.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Index Router is working",
				"status":  "ok",
			})
		})

		index.GET("/api/version", indexController.GetVersion)
	}
}
