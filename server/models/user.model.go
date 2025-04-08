package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/savvy-bit/gin-react-postgres/utils"
	"github.com/savvy-bit/gin-react-postgres/validations"
	"gorm.io/gorm"
)

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
	GuestUser     UserRole = "guest"
)

// implement scan method for reading from database
func (u *UserRole) Scan(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid user role: %v", value)
	}
	*u = UserRole(str)
	return nil
}

// implement value method for writing to database
func (u *UserRole) Value() (driver.Value, error) {
	return string(*u), nil
}

// check if user role is valid
func (u *UserRole) IsRoleValid() bool {
	switch *u {
	case UserRoleAdmin, UserRoleUser, GuestUser:
		return true
	default:
		return false
	}
}

type User struct {
	UserID           uuid.UUID      `gorm:"type:uuid;primaryKey;unique; not null; index" json:"userID"`
	FullName         string         `gorm:"not null " json:"fullName"`
	Username         string         `gorm:"unique;not null" json:"username"`
	Email            string         `gorm:"unique;not null" json:"email"`
	ProfileImage     string         `gorm:"default:null" json:"profileImage"`
	Gender           string         `gorm:"default:null" json:"gender"`
	Role             UserRole       `gorm:"type:user_role;not null;default:'user'" json:"role"`
	BannerImage      string         `gorm:"default:null" json:"bannerImage"`
	PassWord         string         `gorm:"not null" json:"password"`
	AuthOtp          int            `gorm:"default:null" json:"authOtp"`
	IsEmailVerified  bool           `gorm:"default:false" json:"isEmailVerified"`
	ResetPasswordOtp string         `gorm:"default:null" json:"resetPasswordOtp"`
	RefreshToken     string         `gorm:"default:null" json:"refreshToken"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func CreateEnumUserRole(db *gorm.DB) error {
	var exists bool
	err := db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role')").Scan(&exists).Error
	if err != nil {
		return err
	}

	if !exists {
		if err := db.Exec("CREATE TYPE user_role AS ENUM ('admin', 'user', 'guest')").Error; err != nil {
			return err
		}
	}

	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	id := uuid.Must(uuid.NewV4())
	if id != uuid.Nil {
		u.UserID = id
	}
	if err := validations.ValidateUser(validations.User{
		UserID:   u.UserID,
		FullName: u.FullName,
		Username: u.Username,
		Role:     validations.UserRole(u.Role),
		Gender:   validations.UserGender(u.Gender),
		Email:    u.Email,
		Password: u.PassWord,
	}); err != nil {
		return err
	}

	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.PassWord != "" && !strings.HasPrefix(u.PassWord, "$2a$") {
		fmt.Println("Hashing password:", u.PassWord)
		hashedPassword, err := utils.HashPassword(u.PassWord)
		if err != nil {
			return err
		}
		u.PassWord = hashedPassword
	} else {
		fmt.Println("Password already hashed or empty, not hashing again")
	}
	return nil
}

func (User) TableName() string {
	return "users"
}
