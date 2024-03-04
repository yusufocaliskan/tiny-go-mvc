package authmodel

type AuthRefreshTokenData struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
