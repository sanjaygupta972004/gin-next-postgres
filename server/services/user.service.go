package services

import (
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/savvy-bit/gin-react-postgres/config"
	"github.com/savvy-bit/gin-react-postgres/dto"
	"github.com/savvy-bit/gin-react-postgres/models"
	"github.com/savvy-bit/gin-react-postgres/notification/email"
	"github.com/savvy-bit/gin-react-postgres/repositories"
	"github.com/savvy-bit/gin-react-postgres/utils"
	"github.com/savvy-bit/gin-react-postgres/utils/authHelper"
	"github.com/savvy-bit/gin-react-postgres/utils/mapper"
	"github.com/savvy-bit/gin-react-postgres/validations"
)

type UserService interface {
	CreateUser(user *models.User) (*dto.UserResponse, error)
	LoginUser(userLoginReq *dto.UserLoginRequest) (*dto.UserLoginResponse, error)
	LogoutUser(userID string) (message string, err error)
	VerifyEmail(userID string, authOtp int) (message string, err error)
	RegenerateAuthTokens(userID string) (*dto.UserLoginResponse, error)
	RegenerateAuthOtp(userID string) (message string, err error)
	GetUserProfile(userID string) (*dto.UserResponse, error)
	UpdateUserProfile(userID string, updateUserRequest dto.UserUpdateRequest) (*dto.UserResponse, error)
	DeleteUserProfile(userID string) (message string, err error)
	UploadProfileImage(userID string, fileHeader *multipart.FileHeader) (*dto.UserResponse, error)
	UploadBannerImage(userID string, fileHeader *multipart.FileHeader) (*dto.UserResponse, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(user *models.User) (*dto.UserResponse, error) {
	userData, db, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	// send verification email
	otp, err := authHelper.GenerateOTP(6)
	if err != nil {
		return nil, err
	}

	emailBody, err := email.RanderEmailAuthTemplate("email_verification.html", map[string]string{
		"OTP_CODE": otp,
	})

	if err != nil {
		return nil, err
	}

	sslClient, err := config.NewSESClient()
	if err != nil {
		return nil, err
	}

	if err := email.SendEmail(sslClient, user.Email, "Verify your registration", emailBody); err != nil {
		return nil, err
	}
	// update user with OTP and expiry time
	// set expiry time to 1 hour
	currentTime := time.Now()
	if err := db.Model(userData).Updates(map[string]any{
		"auth_otp":             otp,
		"auth_otp_expiry_time": currentTime.Add(1 * time.Hour),
	}).Error; err != nil {
		return nil, err
	}
	return mapper.UserToUserResponse(*userData), nil
}

func (s *userService) RegenerateAuthTokens(userID string) (*dto.UserLoginResponse, error) {
	userUUID, err := utils.IsUUID(userID)
	if err != nil {
		return nil, err
	}
	userData, db, err := s.repo.RegenerateAuthTokens(userUUID)
	if err != nil {
		return nil, err
	}
	accessToken, err := authHelper.SignAccessToken(userData)
	if err != nil {
		return nil, err
	}
	refreshToken, err := authHelper.SignRefreshToken(userData)
	if err != nil {
		return nil, err
	}
	if !userData.IsEmailVerified {
		return nil, errors.New("email not verified")
	}

	if userData.RefreshTokenExpiryTime.Before(time.Now()) {
		return nil, errors.New("refresh token expired please login again")
	}
	if err := db.Model(userData).Updates(map[string]any{
		"refresh_token":             refreshToken,
		"refresh_token_expiry_time": time.Now().Add(7 * 24 * time.Hour),
	}).Error; err != nil {
		return nil, err
	}

	responseUserData := mapper.UserToUserResponse(*userData)
	userLoginResponse := &dto.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Data:         responseUserData,
	}

	return userLoginResponse, nil

}
func (s *userService) VerifyEmail(userID string, authOtp int) (string, error) {
	serverConfig := config.GetGlobalConfig().Server
	if serverConfig.ClientURL == "" {
		return "", errors.New("client URL not configured")
	}

	userUUID, err := utils.IsUUID(userID)
	if err != nil {
		return "", err
	}
	userData, db, err := s.repo.VerifyAuthOtp(userUUID)
	if err != nil {
		return "", err
	}
	if userData.AuthOtp != authOtp {
		return "", errors.New("invalid OTP")
	}
	if time.Now().After(userData.AuthOtpExpiryTime) {
		return "", errors.New("OTP expired")
	}
	if err := db.Model(userData).Updates(map[string]any{
		"is_email_verified":    true,
		"auth_otp":             nil,
		"auth_otp_expiry_time": nil,
	}).Error; err != nil {
		return "", err
	}

	sslClient, err := config.NewSESClient()
	if err != nil {
		return "", err
	}

	emailBody, err := email.RanderWelcomeTemplate("welcome.html", map[string]string{
		"USERNAME": userData.FullName,
		"Href":     serverConfig.ClientURL,
	})
	if err != nil {
		return "", err
	}

	if err := email.SendEmail(sslClient, userData.Email, "Welcome to gin-next ", emailBody); err != nil {
		return "", err
	}

	return "Email verified successfully", nil

}

func (s *userService) RegenerateAuthOtp(userID string) (string, error) {
	userUUID, err := utils.IsUUID(userID)
	if err != nil {
		return "", err
	}
	userData, db, err := s.repo.RegenerateAuthOtp(userUUID)
	if err != nil {
		return "", err
	}

	otp, err := authHelper.GenerateOTP(6)
	if err != nil {
		return "", err
	}

	emailBody, err := email.RanderEmailAuthTemplate("email_verification.html", map[string]string{
		"OTP_CODE": otp,
	})

	if err != nil {
		return "", err
	}

	sslClient, err := config.NewSESClient()
	if err != nil {
		return "", err
	}

	if err := email.SendEmail(sslClient, userData.Email, "Verify your registration", emailBody); err != nil {
		return "", err
	}
	fmt.Println("OTP sent to email:", userData.Email)
	currentTime := time.Now()
	if err := db.Model(userData).Updates(map[string]any{
		"auth_otp":             otp,
		"auth_otp_expiry_time": currentTime.Add(1 * time.Hour),
	}).Error; err != nil {
		return "", err
	}
	return "OTP sent successfully", nil
}

func (s *userService) LoginUser(userLoginReq *dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	userData, db, err := s.repo.LoginUser(*userLoginReq)
	if err != nil {
		return nil, err
	}
	if !userData.IsEmailVerified {
		return &dto.UserLoginResponse{
			AccessToken:  "",
			RefreshToken: "",
			Data:         mapper.UserToUserResponse(*userData),
		}, nil
	}

	if err := utils.CompareHashAndPassword(userData.PassWord, userLoginReq.Password); err != nil {
		return nil, err
	}
	accessToken, err := authHelper.SignAccessToken(userData)
	if err != nil {
		return nil, err
	}
	refreshToken, err := authHelper.SignRefreshToken(userData)
	if err != nil {
		return nil, err
	}

	if err := db.Model(userData).Updates(map[string]any{
		"refresh_token":             refreshToken,
		"refresh_token_expiry_time": time.Now().Add(7 * 24 * time.Hour),
	}).Error; err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	responseUserData := mapper.UserToUserResponse(*userData)

	userLoginResponse := &dto.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Data:         responseUserData,
	}
	return userLoginResponse, nil
}

func (s *userService) GetUserProfile(userID string) (*dto.UserResponse, error) {
	userUUID, err := utils.IsUUID(userID)
	if err != nil {
		return nil, err
	}
	userData, err := s.repo.GetUserByID(userUUID)
	if err != nil {
		return nil, err
	}
	return mapper.UserToUserResponse(*userData), nil
}

func (s *userService) LogoutUser(userID string) (string, error) {
	userUUID, err := utils.IsUUID(userID)
	if err != nil {
		return "", err
	}

	if err := s.repo.LogoutUser(userUUID); err != nil {
		return "", err
	}
	return "Logout successfully", nil
}

func (s *userService) UpdateUserProfile(userID string, updateUserRequest dto.UserUpdateRequest) (*dto.UserResponse, error) {
	userUUID, err := utils.IsUUID(userID)
	if err != nil {
		return nil, err
	}
	gender := validations.UserGender(updateUserRequest.Gender)
	isValidGender := validations.IsValidGender(&gender)
	if !isValidGender {
		return nil, errors.New("invalid user gender")
	}
	userData, err := s.repo.UpdateUser(userUUID, dto.UserUpdateRequest{
		FullName: updateUserRequest.FullName,
		Username: updateUserRequest.Username,
		Gender:   string(gender),
	})
	if err != nil {
		return nil, err
	}
	return mapper.UserToUserResponse(*userData), nil
}

func (s *userService) DeleteUserProfile(userID string) (string, error) {
	userUUID, err := utils.IsUUID(userID)
	if err != nil {
		return "", err
	}
	if err := s.repo.DeleteUser(userUUID); err != nil {
		return "", err
	}
	return "User deleted successfully", nil
}

func (s *userService) UploadProfileImage(userID string, fileHeader *multipart.FileHeader) (*dto.UserResponse, error) {
	userUUID, err := utils.IsUUID(userID)
	if err != nil {
		return nil, err
	}

	if fileHeader == nil {
		return nil, errors.New("file is required")
	}

	// Check file size, maximum size is 5MB
	if fileHeader.Size > 5*1024*1024 {
		return nil, errors.New("file size should not exceed 5MB")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	profileImageURL, err := utils.UploadFileToS3(file, fileHeader)
	if err != nil {
		return nil, err
	}

	if profileImageURL == "" {
		return nil, errors.New("failed to upload profile image")
	}
	userData, err := s.repo.UploadProfileImage(userUUID, profileImageURL)

	if err != nil {
		return nil, err
	}

	fmt.Println("Profile image uploaded successfully", userData)
	return mapper.UserToUserResponse(*userData), nil
}

func (s *userService) UploadBannerImage(userID string, fileHeader *multipart.FileHeader) (*dto.UserResponse, error) {
	userUUID, err := utils.IsUUID(userID)
	if err != nil {
		return nil, err
	}

	if fileHeader == nil {
		return nil, errors.New("file is required")
	}

	// Check file size, maximum size is 5MB
	if fileHeader.Size > 5*1024*1024 {
		return nil, errors.New("file size should not exceed 5MB")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bannerImageURL, err := utils.UploadFileToS3(file, fileHeader)
	if err != nil {
		return nil, err
	}

	if bannerImageURL == "" {
		return nil, errors.New("failed to upload banner image")
	}
	userData, err := s.repo.UploadBannerImage(userUUID, bannerImageURL)

	if err != nil {
		return nil, err
	}
	return mapper.UserToUserResponse(*userData), nil
}
