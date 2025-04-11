package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/controllers"
	"github.com/savvy-bit/gin-react-postgres/middlewares"
	"github.com/savvy-bit/gin-react-postgres/repositories"
	"github.com/savvy-bit/gin-react-postgres/services"
	"gorm.io/gorm"
)

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
		user.POST("/regenerate-auth-otp/:userID", userController.RegenerateAuthOtp)
		user.GET("/logout", authMiddleware, userController.LogoutUser)
		user.GET("/regenerate-auth-tokens", userController.RegenerateAuthTokens)
		user.GET("/get-profile", authMiddleware, userController.GetUserProfile)

	}
}
