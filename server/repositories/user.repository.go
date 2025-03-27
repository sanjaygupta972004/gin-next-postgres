package repositories

import (
	"github.com/gofrs/uuid"
	"github.com/savvy-bit/gin-react-postgres/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	UploadProfileImage(userID uuid.UUID, profileImage string) (*models.User, error)
	UploadBannerImage(userID uuid.UUID, bannerImage string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	LoginUser(email string) (*models.User, error)
	LogoutUser() (message string, err error)
	GetUserByID(userID uuid.UUID) (*models.User, error)
	UpateUser(userID uuid.UUID, user *models.User) (*models.User, error)
	DeleteUser(userID uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// CreateUser implements UserRepository.
func (u *userRepository) CreateUser(user *models.User) (*models.User, error) {
	panic("unimplemented")
}

// LoginUser implements UserRepository.
func (u *userRepository) LoginUser(email string) (*models.User, error) {
	panic("unimplemented")
}

// LogoutUser implements UserRepository.
func (u *userRepository) LogoutUser() (string, error) {
	panic("unimplemented")
}

// DeleteUser implements UserRepository.
func (u *userRepository) DeleteUser(userID uuid.UUID) error {
	panic("unimplemented")
}

// GetUserByID implements UserRepository.
func (u *userRepository) GetUserByID(userID uuid.UUID) (*models.User, error) {
	panic("unimplemented")
}

// UpateUser implements UserRepository.
func (u *userRepository) UpateUser(userID uuid.UUID, user *models.User) (*models.User, error) {
	panic("unimplemented")
}

// UploadBannerImage implements UserRepository.
func (u *userRepository) UploadBannerImage(userID uuid.UUID, bannerImage string) (*models.User, error) {
	panic("unimplemented")
}

// UploadProfileImage implements UserRepository.
func (u *userRepository) UploadProfileImage(userID uuid.UUID, profileImage string) (*models.User, error) {
	panic("unimplemented")
}
