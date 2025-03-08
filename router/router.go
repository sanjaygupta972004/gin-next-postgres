package router

import (
	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/middleware"
	"github.com/savvy-bit/gin-react-postgres/model"
)

func Route(app *gin.Engine) {
	authMiddleware := middleware.Auth()

	auth := app.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	// Endpoints need authentication
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			// claims := jwt.ExtractClaims(c)
			user, _ := c.Get("email")
			c.JSON(200, gin.H{
				// "email": claims["email"],
				"name": user.(*model.User).Name,
				"text": "Hello World.",
			})
		})
	}

	// Auth endpoints
	app.POST(
		"/login", authMiddleware.LoginHandler,
	)
}
