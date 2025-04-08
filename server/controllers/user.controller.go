package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/dto"
	"github.com/savvy-bit/gin-react-postgres/models"
	"github.com/savvy-bit/gin-react-postgres/services"
	"github.com/savvy-bit/gin-react-postgres/utils"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	LogoutUser(c *gin.Context)
	VerifyEmail(c *gin.Context)
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

func (u *userController) RegisterUser(c *gin.Context) {
	var userRegisterRequest dto.UserRegisterRequest

	if err := utils.ValidateJSONBody(c, &userRegisterRequest); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to validate JSON body", utils.ErrBadRequest)
		return
	}
	newUser := &models.User{
		FullName: userRegisterRequest.FullName,
		Username: userRegisterRequest.Username,
		Email:    userRegisterRequest.Email,
		Gender:   userRegisterRequest.Gender,
		PassWord: userRegisterRequest.Password,
		Role:     models.UserRole(userRegisterRequest.Role),
	}

	user, err := u.userService.CreateUser(newUser)

	if err != nil {
		fmt.Println("Error creating user:", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to register user", utils.ErrInternalServer)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "User registered successfully", user)
}

func (u *userController) VerifyEmail(c *gin.Context) {
	userID := c.Param("userID")
	var verifyEmailRequest struct {
		AuthOtp int `json:"authOtp" binding:"required"`
	}
	if userID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "User ID is required", utils.ErrBadRequest)
		return
	}

	if err := utils.ValidateJSONBody(c, &verifyEmailRequest); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to validate JSON body", utils.ErrBadRequest)
		return
	}

	message, err := u.userService.VerifyEmail(userID, verifyEmailRequest.AuthOtp)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to verify email: %v", err), utils.ErrInternalServer)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, message, nil)
}

func (u *userController) LoginUser(c *gin.Context) {
	var userLoginRequest dto.UserLoginRequest

	if err := utils.ValidateJSONBody(c, &userLoginRequest); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to validate JSON body", utils.ErrBadRequest)
		return
	}
	user, err := u.userService.LoginUser(&userLoginRequest)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("failed to login user: %v", err), utils.ErrUnauthorized)
		return
	}
	c.SetCookie("accessToken", user.AccessToken, 3600, "/", "", false, true)
	c.SetCookie("refreshToken", user.RefreshToken, 3600*24*7, "/", "", false, true)
	utils.SuccessResponse(c, http.StatusOK, "User logged in successfully", user)
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
