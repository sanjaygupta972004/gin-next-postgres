package services

import (
	"errors"

	"github.com/savvy-bit/gin-react-postgres/dto"
	"github.com/savvy-bit/gin-react-postgres/models"
	"github.com/savvy-bit/gin-react-postgres/repositories"
	"github.com/savvy-bit/gin-react-postgres/utils"
	"github.com/savvy-bit/gin-react-postgres/utils/mapper"
	"github.com/savvy-bit/gin-react-postgres/validations"
)

type UserService interface {
	CreateUser(user *models.User) (*dto.UserResponse, error)
	LoginUser(user *models.User) (*dto.UserLoginResponse, error)
	LogoutUser(userID string) (message string, err error)
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
	userData, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return mapper.UserToUserResponse(*userData), nil
}

// will have to implement
func (s *userService) LoginUser(user *models.User) (*dto.UserLoginResponse, error) {
	panic("unimplemented")
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
