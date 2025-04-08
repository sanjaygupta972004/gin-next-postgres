package dto

type UserUpdateRequest struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserRegisterRequest struct {
	FullName string `json:"fullName" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Gender   string `json:"gender" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=admin user guest"`
}

type UserResponse struct {
	UserID       string `json:"userID"`
	FullName     string `json:"fullName"`
	Gender       string `json:"gender"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	Username     string `json:"username"`
	ProfileImage string `json:"profileImage"`
	BannerImage  string `json:"bannerImage"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
	DeletedAt    string `json:"deletedAt,omitempty"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
