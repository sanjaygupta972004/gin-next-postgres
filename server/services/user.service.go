package services

import (
	"errors"

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
	GetUserProfile(userID string) (*dto.UserResponse, error)
	UpdateUserProfile(userID string, updateUserRequest dto.UserUpdateRequest) (*dto.UserResponse, error)
	DeleteUserProfile(userID string) (message string, err error)
	UploadProfileImage(userID string, profileImage string) (*dto.UserResponse, error)
	UploadBannerImage(userID string, bannerImage string) (*dto.UserResponse, error)
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

	sslClient, err := email.SESAWSClient()
	if err != nil {
		return nil, err
	}

	if err := email.SendEmail(sslClient, user.Email, "Verify your registration", emailBody); err != nil {
		return nil, err
	}
	if err := db.Model(userData).Update("auth_otp", otp).Error; err != nil {
		return nil, err
	}

	return mapper.UserToUserResponse(*userData), nil
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
	userData, err := s.repo.VerifyAuthOtp(userUUID, authOtp)
	if err != nil {
		return "", err
	}
	sslClient, err := email.SESAWSClient()
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

func (s *userService) LoginUser(userLoginReq *dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	userData, db, err := s.repo.LoginUser(*userLoginReq)
	if err != nil {
		return nil, err
	}
	if !userData.IsEmailVerified {
		return nil, errors.New("email not verified")
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

	if err := db.Model(userData).Update("refresh_token", refreshToken).Error; err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	userLoginResponse := &dto.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
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
	userData, err := s.repo.UpdateUser(userUUID, &models.User{
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

// UploadProfileImage
func (s *userService) UploadProfileImage(userID string, profileImage string) (*dto.UserResponse, error) {
	panic("unimplemented")
}

// uploadBannerImage
func (s *userService) UploadBannerImage(userID string, bannerImage string) (*dto.UserResponse, error) {
	panic("unimplemented")
}
