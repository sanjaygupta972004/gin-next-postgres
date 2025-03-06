package router

import (
	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/controller"
)

func Route(app *gin.Engine) {
	indexController := new(controller.IndexController)
	userController := new(controller.UserController)

	app.GET(
		"/", indexController.GetIndex,
	)

	user := app.Group("/user")
	{
		user.POST("/signup", userController.SignUp)
	}
	api := app.Group("/api")
	{
		api.GET("/version", indexController.GetVersion)
	}
}
