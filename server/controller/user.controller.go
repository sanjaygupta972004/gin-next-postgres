package controller

import (
	"net/http"

	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/model"
)

// UserController is the default controller
type UserController struct{}

type SignUpRequest struct {
	Email     string `form:"email" json:"email" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required,min=6,max=20"`
	Password2 string `form:"password2" json:"password2" binding:"required"`
}

type GetMeResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// @Summary Get User Information
// @Description This endpoint returns the user information
// @Tags Auth
// @Accept json
// @Produce json
// @Security  Bearer
// @Success 200 {object} GetMeResponse "Successful response"
// @Router /auth/me [get]
func (ctrl *UserController) GetMe(c *gin.Context) {
	claims := ginjwt.ExtractClaims(c)
	c.JSON(http.StatusOK, gin.H{
		"email": claims["email"].(string),
		"name":  claims["name"].(string),
		"role":  claims["role"].(string),
	})
}

func (ctrl *UserController) SignUp(c *gin.Context) {
	var request SignUpRequest
	if err := c.ShouldBind(&request); err == nil {

		if request.Password != request.Password2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password does not match with conform password"})
			return
		}

		var user model.User

		user.Name = request.Name
		user.Email = request.Email
		user.Password = request.Password

		if err := user.Signup(); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, user)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := model.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
