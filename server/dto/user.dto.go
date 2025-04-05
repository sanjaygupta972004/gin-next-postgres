package dto

type UserUpdateRequest struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
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
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
