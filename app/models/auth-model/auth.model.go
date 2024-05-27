package authmodel

import (
	"time"

	tinytoken "github.com/gptverse/init/framework/tiny-token"
)

type AuthRefreshTokenModel struct {
	RefreshToken string `json:"refresh_token" validate:"required" example:"the_refresh_token"`
}

type AuthLoginModel struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}
type AuthUserTokenModel struct {
	Email     string                  `bson:"email" json:"email"`
	Token     tinytoken.TinyTokenData `bson:"token" `
	Status    string                  `bson:"status" `
	UpdatedAt time.Time               `bson:"updated_at" `
}

type AuthModel struct {
	tinytoken.TinyToken
}

type SessionModel struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}
