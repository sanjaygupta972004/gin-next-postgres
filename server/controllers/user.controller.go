package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/services"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	LogoutUser(c *gin.Context)

	GetUserProfile(c *gin.Context)
	UploadBannerImage(c *gin.Context)
	UploadProfileImage(c *gin.Context)
	DeleteUserProfile(c *gin.Context)
	UpdateUserProfile(c *gin.Context)
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

type userController struct {
	userService services.UserService
}

// DeleteUserProfile implements UserController.
func (u *userController) DeleteUserProfile(c *gin.Context) {
	panic("unimplemented")
}

// GetUserProfile implements UserController.
func (u *userController) GetUserProfile(c *gin.Context) {
	panic("unimplemented")
}

// LoginUser implements UserController.
func (u *userController) LoginUser(c *gin.Context) {
	panic("unimplemented")
}

// LogoutUser implements UserController.
func (u *userController) LogoutUser(c *gin.Context) {
	panic("unimplemented")
}

// RegisterUser implements UserController.
func (u *userController) RegisterUser(c *gin.Context) {
	panic("unimplemented")
}

// UpdateUserProfile implements UserController.
func (u *userController) UpdateUserProfile(c *gin.Context) {
	panic("unimplemented")
}

// UploadBannerImage implements UserController.
func (u *userController) UploadBannerImage(c *gin.Context) {
	panic("unimplemented")
}

// UploadProfileImage implements UserController.
func (u *userController) UploadProfileImage(c *gin.Context) {
	panic("unimplemented")
}
