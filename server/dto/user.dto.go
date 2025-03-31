package dto

import "github.com/savvy-bit/gin-react-postgres/models"

type UserUpdateRequest struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Role     string `json:"role"`
	IsAdmin  bool   `json:"isAdmin"`
	Gender   string `json:"gender"`
}

type UserResponse struct {
	UserID       string          `json:"userID"`
	FullName     string          `json:"fullName"`
	Gender       string          `json:"gender"`
	Email        string          `json:"email"`
	Role         models.UserRole `json:"role"`
	Username     string          `json:"username"`
	ProfileImage string          `json:"profileImage"`
	BannerImage  string          `json:"bannerImage"`
	IsAdmin      bool            `json:"isAdmin"`
	CreatedAt    string          `json:"createdAt"`
	UpdatedAt    string          `json:"updatedAt"`
	DeletedAt    string          `json:"deletedAt,omitempty"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
