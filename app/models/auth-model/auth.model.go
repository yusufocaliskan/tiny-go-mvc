package authmodel

import tinytoken "github.com/gptverse/init/framework/tiny-token"

type AuthRefreshTokenModel struct {
	RefreshToken string `json:"refresh_token" validate:"required" example:"the_refresh_token"`
}

type AuthModel struct {
	tinytoken.TinyToken
}
