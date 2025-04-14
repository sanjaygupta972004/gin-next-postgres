package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/savvy-bit/gin-react-postgres/config"
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
	RegenerateAuthOtp(c *gin.Context)
	RegenerateAuthTokens(c *gin.Context)
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
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %v", err), utils.ErrInternalServer)
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
func (u *userController) RegenerateAuthOtp(c *gin.Context) {
	userID := c.Param("userID")
	if userID == "" {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to get user ID from path:", utils.ErrInternalServer)
		return
	}
	message, err := u.userService.RegenerateAuthOtp(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to regenerate auth OTP: %v", err), utils.ErrInternalServer)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, message, nil)
}

func (u *userController) RegenerateAuthTokens(c *gin.Context) {
	var token string

	authConfig := config.GetGlobalConfig().AuthToken

	if cookieToken, err := c.Cookie("refreshToken"); err == nil {
		token = cookieToken
	}

	if token == "" {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			token = strings.TrimPrefix(strings.TrimSpace(authHeader), "Bearer ")
		}
	}

	if token == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Token not provided in Authorization header or cookie", utils.ErrBadRequest)
		return
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return []byte(authConfig.RefreshToken), nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized access", utils.ErrUnauthorized)
		return
	}

	log.Println("Claims:", claims)

	userID, ok := claims["userID"].(string)
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token: missing or invalid id", utils.ErrUnauthorized)
		return
	}

	userData, err := u.userService.RegenerateAuthTokens(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to regenerate auth tokens: %v", err), utils.ErrInternalServer)
		return
	}
	c.SetCookie("accessToken", userData.AccessToken, 3600*12, "/", "", false, true)
	c.SetCookie("refreshToken", userData.RefreshToken, 3600*24*7, "/", "", false, true)
	utils.SuccessResponse(c, http.StatusOK, "Auth tokens regenerated successfully", userData)
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

	if !user.Data.IsEmailVerified {
		utils.SuccessResponse(c, http.StatusOK, "User Email is not verified", user)
		return
	}
	c.SetCookie("accessToken", user.AccessToken, 3600*12, "/", "", false, true)
	c.SetCookie("refreshToken", user.RefreshToken, 3600*24*7, "/", "", false, true)
	utils.SuccessResponse(c, http.StatusOK, "User logged in successfully", user)
}

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

func (u *userController) DeleteUserProfile(c *gin.Context) {
	userID, err := utils.GetUserIdFromHeader(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("failed to get user ID from header: %v", err), utils.ErrUnauthorized)
		return
	}
	message, err := u.userService.DeleteUserProfile(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to delete user profile: %v", err), utils.ErrInternalServer)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, message, nil)
}

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

func (u *userController) UploadBannerImage(c *gin.Context) {
	userID, err := utils.GetUserIdFromHeader(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("failed to get user ID from header: %v", err), utils.ErrUnauthorized)
		return
	}
	file, err := c.FormFile("bannerImage")
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to get form file: %v", err), utils.ErrBadRequest)
		return
	}

	userUpdatedData, err := u.userService.UploadBannerImage(userID, file)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to upload banner image: %v", err), utils.ErrInternalServer)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Banner image uploaded successfully", userUpdatedData)

}

func (u *userController) UploadProfileImage(c *gin.Context) {
	userID, err := utils.GetUserIdFromHeader(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("failed to get user ID from header: %v", err), utils.ErrUnauthorized)
		return
	}
	file, err := c.FormFile("profileImage")
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to get form file: %v", err), utils.ErrBadRequest)
		return
	}

	userUpdatedData, err := u.userService.UploadProfileImage(userID, file)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to upload profile image: %v", err), utils.ErrInternalServer)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Profile image uploaded successfully", userUpdatedData)

}
