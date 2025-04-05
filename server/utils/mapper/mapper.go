package mapper

import (
	"github.com/savvy-bit/gin-react-postgres/dto"
	"github.com/savvy-bit/gin-react-postgres/models"
)

func UserToUserResponse(user models.User) *dto.UserResponse {
	return &dto.UserResponse{
		UserID:    user.UserID.String(),
		FullName:  user.FullName,
		Username:  user.Username,
		Email:     user.Email,
		Role:      string(user.Role),
		Gender:    user.Gender,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		DeletedAt: func() string {
			if user.DeletedAt.Valid {
				return user.DeletedAt.Time.String()
			}
			return ""
		}(),
		ProfileImage: func() string {
			if user.ProfileImage != "" {
				return user.ProfileImage
			}
			return ""
		}(),
		BannerImage: func() string {
			if user.BannerImage != "" {
				return user.BannerImage
			}
			return ""
		}(),
	}
}
