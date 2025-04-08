package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/controllers"
	"github.com/savvy-bit/gin-react-postgres/middlewares"
	"github.com/savvy-bit/gin-react-postgres/repositories"
	"github.com/savvy-bit/gin-react-postgres/services"
	"gorm.io/gorm"
)

// func restrictToRoles(allowedRoles []string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		claims := ginjwt.ExtractClaims(c)

// 		if claims["role"] == nil {
// 			c.AbortWithStatusJSON(401, gin.H{"error": "User not found"})
// 			return
// 		}
// 		userRole := claims["role"]
// 		for _, role := range allowedRoles {
// 			if userRole == role {
// 				c.Next()
// 				return
// 			}
// 		}
// 		c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
// 	}
// }

// func Route(app *gin.Engine) {
// 	indexController := new(controllers.IndexController)
// 	userController := new(controllers.UserController)
// 	authMiddleware := middleware.Auth()

// 	// Public endpoints
// 	app.POST("/login", authMiddleware.LoginHandler)

// 	// Admin endpoints
// 	admin := app.Group("/admin")
// 	admin.Use(authMiddleware.MiddlewareFunc())
// 	{
// 		admin.GET("/users", restrictToRoles([]string{"admin"}), userController.GetAllUsers)
// 	}

// 	// Auth endpoints
// 	auth := app.Group("/auth")
// 	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
// 	auth.Use(authMiddleware.MiddlewareFunc())
// 	{
// 		auth.GET("/me", userController.GetMe)
// 	}

// 	// Api
// 	api := app.Group("/api")
// 	api.Use()
// 	{
// 		api.GET("/version", indexController.GetVersion)
// 	}
// }

func SetUpUserRouter(router *gin.RouterGroup, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	userController := controllers.NewUserController(userService)
	authMiddleware := middlewares.JWTVerifyForUser(db)

	user := router.Group("/user")
	{
		user.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "User Router is working",
				"status":  "ok",
			})
		})

		user.POST("/register", userController.RegisterUser)
		user.POST(("/verify-email/:userID"), userController.VerifyEmail)
		user.POST("/login", userController.LoginUser)
		user.GET("/get-profile", authMiddleware, userController.GetUserProfile)

	}
}
