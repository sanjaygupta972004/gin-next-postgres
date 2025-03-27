package services

import (
	"github.com/savvy-bit/gin-react-postgres/models"
	"github.com/savvy-bit/gin-react-postgres/repositories"
)

type LoginResponse struct {
	Token string `json:"token"`
}

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	LoginUser(user *models.User) (*LoginResponse, error)
	LogoutUser() (message string, err error)
	GetUserProfile(userID string) (*models.User, error)
	UpdateUserProfile(userID string, user *models.User) (*models.User, error)
	DeleteUserProfile(userID string) error
	UploadProfileImage(userID string, profileImage string) (*models.User, error)
	UploadBannerImage(userID string, bannerImage string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// DeleteUserProfile implements UserService.
func (s *userService) DeleteUserProfile(userID string) error {
	panic("unimplemented")
}

// GetUserProfile implements UserService.
func (s *userService) GetUserProfile(userID string) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUserProfile implements UserService.
func (s *userService) UpdateUserProfile(userID string, user *models.User) (*models.User, error) {
	panic("unimplemented")
}

// UploadBannerImage implements UserService.
func (s *userService) UploadBannerImage(userID string, bannerImage string) (*models.User, error) {
	panic("unimplemented")
}

// UploadProfileImage implements UserService.
func (s *userService) UploadProfileImage(userID string, profileImage string) (*models.User, error) {
	panic("unimplemented")
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	return s.repo.CreateUser(user)
}

func (s *userService) LoginUser(user *models.User) (*LoginResponse, error) {
	panic("unimplemented")
}

func (s *userService) LogoutUser() (string, error) {
	panic("unimplemented")
}
