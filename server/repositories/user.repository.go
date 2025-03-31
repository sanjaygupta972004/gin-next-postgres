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
	LogoutUser(userID uuid.UUID) error
	GetUserByID(userID uuid.UUID) (*models.User, error)
	UpdateUser(userID uuid.UUID, user *models.User) (*models.User, error)
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

func (u *userRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := user.BeforeCreate(u.db); err != nil {
		return nil, err
	}
	if err := u.db.Find("email", user.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := u.db.Create(&user).Error; err != nil {
				return nil, err
			}
			return user, nil
		}
	}
	return nil, nil
}

// LoginUser implements UserRepository.
func (u *userRepository) LoginUser(email string) (*models.User, error) {
	panic("unimplemented")
}

// LogoutUser implements UserRepository.
func (u *userRepository) LogoutUser(userID uuid.UUID) error {
	var user models.User
	if err := u.db.Model(&user).Where("user_id = ?", userID).Updates(map[string]interface{}{
		"refresh_token": nil,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) DeleteUser(userID uuid.UUID) error {
	var user models.User
	if err := u.db.Delete(&user, "user_id = ?", userID).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := u.db.First(&user, "user_id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) UpdateUser(userID uuid.UUID, user *models.User) (*models.User, error) {
	if err := u.db.Model(&user).Where("user_id = ?", userID).Updates(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// UploadBannerImage implements UserRepository.
func (u *userRepository) UploadBannerImage(userID uuid.UUID, bannerImage string) (*models.User, error) {
	panic("unimplemented")
}

// UploadProfileImage implements UserRepository.
func (u *userRepository) UploadProfileImage(userID uuid.UUID, profileImage string) (*models.User, error) {
	panic("unimplemented")
}
