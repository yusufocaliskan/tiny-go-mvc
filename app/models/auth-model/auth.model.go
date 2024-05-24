package authmodel

import (
	tinytoken "github.com/gptverse/init/framework/tiny-token"
)

type AuthRefreshTokenModel struct {
	RefreshToken string `json:"refresh_token" validate:"required" example:"the_refresh_token"`
}

type AuthLoginModel struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

type AuthModel struct {
	tinytoken.TinyToken
}
