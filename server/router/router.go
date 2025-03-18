package router

import (
	"fmt"

	ginjwt "github.com/appleboy/gin-jwt/v2"
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
	user, _ := c.Get("email")
	c.JSON(200, gin.H{
		"name": user.(*model.User).Name,
		"text": fmt.Sprintf("Hello %v! Welcome to the Gin + Postgres world.", user.(*model.User).Name),
	})
}

func restrictToRoles(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := ginjwt.ExtractClaims(c)

		if claims["role"] == nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "User not found"})
			return
		}
		userRole := claims["role"]
		for _, role := range allowedRoles {
			if userRole == role {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
	}
}

func Route(app *gin.Engine) {
	indexController := new(controller.IndexController)
	userController := new(controller.UserController)
	authMiddleware := middleware.Auth()

	// Auth endpoints
	app.POST(
		"/login", authMiddleware.LoginHandler,
	)

	auth := app.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", restrictToRoles([]string{"admin"}), getHello)
	}

	app.Use(authMiddleware.MiddlewareFunc())
	{
		app.GET("/users", userController.GetAllUsers)
	}

	api := app.Group("/api")
	api.Use()
	{
		api.GET("/version", indexController.GetVersion)
	}
}
