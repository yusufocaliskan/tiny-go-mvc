package authmodel

type AuthRefreshTokenModel struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
