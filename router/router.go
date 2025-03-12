package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/controller"
	"github.com/savvy-bit/gin-react-postgres/middleware"
	"github.com/savvy-bit/gin-react-postgres/model"
)

// @Summary Test Authorization
// @Description This endpoint is available for only authenticated users
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Security  Bearer
// @Router /auth/hello [get]
func getHello(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	user, _ := c.Get("email")
	c.JSON(200, gin.H{
		// "email": claims["email"],
		"name": user.(*model.User).Name,
		"text": fmt.Sprintf("Hello %v! Welcome to the Gin + Postgres world.", user.(*model.User).Name),
	})
}

func Route(app *gin.Engine) {
	indexController := new(controller.IndexController)
	authMiddleware := middleware.Auth()

	// Auth endpoints
	app.POST(
		"/login", authMiddleware.LoginHandler,
	)

	auth := app.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", getHello)
	}

	api := app.Group("/api")
	api.Use()
	{
		api.GET("/version", indexController.GetVersion)
	}
}
