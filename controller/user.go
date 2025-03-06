package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController is the default controller
type UserController struct{}

// Create new user
type SignUpRequest struct {
	Email string `form:"email" json:"email" binding:"required"`
	Name  string `form:"name" json:"name" binding:"required"`
}

func (ctrl *UserController) SignUp(c *gin.Context) {
	var request SignUpRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": request,
	})
}
