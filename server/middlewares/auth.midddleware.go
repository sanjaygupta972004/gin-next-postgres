package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/savvy-bit/gin-react-postgres/config"
	"github.com/savvy-bit/gin-react-postgres/models"
	"github.com/savvy-bit/gin-react-postgres/utils"
	"gorm.io/gorm"
)

func JWTVerifyForUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		authConfig := config.GetGlobalConfig().AuthToken

		if cookieToken, err := c.Cookie("accessToken"); err == nil {
			token = cookieToken
		}

		if token == "" {
			authHeader := c.GetHeader("Authorization")
			if authHeader != "" {
				token = strings.TrimPrefix(strings.TrimSpace(authHeader), "Bearer ")
			}
		}

		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Token not provided in Authorization header or cookie"})
			c.Abort()
			return
		}

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
			return []byte(authConfig.AccessToken), nil
		})

		if err != nil {
			log.Printf("Error parsing token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		log.Println("Claims:", claims)

		userID, ok := claims["userID"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: missing or invalid id"})
			c.Abort()
			return
		}

		userEmail, ok := claims["email"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: missing or invalid email"})
			c.Abort()
			return
		}

		userRole, ok := claims["role"].(string)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: missing or invalid role"})
			c.Abort()
			return
		}

		if userID == "" || userEmail == "" || userRole != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: missing or invalid userID or email or role"})
			c.Abort()
			return
		}

		log.Printf("User email and userID after authorized access token: %s, %v", userEmail, userID)

		idUUID, err := utils.IsUUID(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: invalid id"})
			c.Abort()
			return
		}

		var user models.User
		result := db.Where("user_id = ? AND email = ?", idUUID, userEmail).First(&user)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
			} else {
				log.Printf("Database error: %v", result.Error)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user in database"})
			}
			c.Abort()
			return
		}

		userStruct := map[string]any{
			"userID": user.UserID.String(),
			"email":  user.Email,
			"role":   user.Role,
		}
		c.Set("user", userStruct)
		c.Next()
	}
}
