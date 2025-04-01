package validations

import (
	"errors"
	"regexp"

	"github.com/gofrs/uuid"
)

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
	GuestUser     UserRole = "guest"
)

type UserGender string

const (
	Male   UserGender = "male"
	Female UserGender = "female"
	Other  UserGender = "other"
)

func IsValidGender(ug *UserGender) bool {
	switch *ug {
	case Male, Female, Other:
		return true
	default:
		return false
	}
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	number := regexp.MustCompile(`[0-9]`)
	special := regexp.MustCompile(`[!@#~$%^&*()+|_]{1}`)
	uppercase := regexp.MustCompile(`[A-Z]`)

	if !number.MatchString(password) {
		return errors.New("password must contain at least one number")
	}

	if !special.MatchString(password) {
		return errors.New("password must contain at least one special character")
	}

	if !uppercase.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	return nil
}

func validateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]{1,64}@[a-zA-Z0-9\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("enter a valid email")
	}
	return nil
}

type User struct {
	UserID   uuid.UUID
	FullName string
	Username string
	Email    string
	Password string
	Gender   UserGender
	Role     UserRole
}

func ValidateUser(u User) error {
	if err := validatePassword(u.Password); err != nil {
		return err
	}
	if err := validateEmail(u.Email); err != nil {
		return err
	}

	if u.UserID == uuid.Nil {
		return errors.New("user id is required")
	}

	if u.FullName == "" {
		return errors.New("full name is required")
	}

	if u.Username == "" {
		return errors.New("username is required")
	}

	if u.Role != UserRoleAdmin && u.Role != UserRoleUser && u.Role != GuestUser {
		return errors.New("invalid user role")
	}
	isValid := IsValidGender(&u.Gender)
	if !isValid {
		return errors.New("invalid user gender")
	}
	return nil
}
