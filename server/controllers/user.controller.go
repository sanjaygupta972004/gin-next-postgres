package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/dto"
	"github.com/savvy-bit/gin-react-postgres/services"
	"github.com/savvy-bit/gin-react-postgres/utils"
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
	userID, err := utils.GetUserIdFromHeader(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("failed to get user ID from header: %v", err), utils.ErrUnauthorized)
		return
	}

	user, err := u.userService.GetUserProfile(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, fmt.Sprintf("failed to get user profile: %v", err), utils.ErrNotFound)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "User profile retrieved successfully", user)
}

// LoginUser implements UserController.
func (u *userController) LoginUser(c *gin.Context) {
	panic("unimplemented")
}

func (u *userController) LogoutUser(c *gin.Context) {
	userID, err := utils.GetUserIdFromHeader(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("failed to get user ID from header: %v", err), utils.ErrUnauthorized)
		return
	}
	message, err := u.userService.LogoutUser(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to logout user: %v", err), utils.ErrInternalServer)
		return
	}
	c.SetCookie("accessToken", "", -1, "/", "", false, true)
	c.SetCookie("refreshToken", "", -1, "/", "", false, true)
	utils.SuccessResponse(c, http.StatusOK, message, nil)
}

// RegisterUser implements UserController.
func (u *userController) RegisterUser(c *gin.Context) {
	panic("unimplemented")
}

// UpdateUserProfile implements UserController.
func (u *userController) UpdateUserProfile(c *gin.Context) {
	userID, err := utils.GetUserIdFromHeader(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("failed to get user ID from header: %v", err), utils.ErrUnauthorized)
		return
	}
	var userUpdateRequest dto.UserUpdateRequest

	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to bind JSON: %v", err), utils.ErrBadRequest)
		return
	}
	user, err := u.userService.UpdateUserProfile(userID, userUpdateRequest)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to update user profile: %v", err), utils.ErrInternalServer)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "User profile updated successfully", user)

}

// UploadBannerImage implements UserController.
func (u *userController) UploadBannerImage(c *gin.Context) {
	panic("unimplemented")
}

// UploadProfileImage implements UserController.
func (u *userController) UploadProfileImage(c *gin.Context) {
	panic("unimplemented")
}
