package tinytoken

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SingleToken struct {
	ExpiryTime time.Duration `json:"expiry_time"`
	BearerKey  string        `json:"brearer_key"`
}
type TinyToken struct {
	AccessToken  SingleToken `json:"access_oken"`
	RefreshToken SingleToken `json:"refresh_token"`
}

var secretKey = []byte("secret-key-test")

func (tToken *TinyToken) Genera(data interface{}) {

	tToken.AccessTokenGenerator(data)
	tToken.RefreshTokenGenerator(data)
}

func (tToken *TinyToken) CreateToken(data interface{}, expiryTime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(expiryTime),
	})

	stringifiedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return stringifiedToken, nil
}

// Creates Refresh token for 1 day
func (tToken *TinyToken) AccessTokenGenerator(data interface{}) {

	expiryTime := time.Hour * 24
	token, _ := tToken.CreateToken(data, expiryTime)
	tToken.AccessToken = SingleToken{
		BearerKey:  token,
		ExpiryTime: expiryTime,
	}
}

// Creates Refresh token for 7 days
func (tToken *TinyToken) RefreshTokenGenerator(data interface{}) {

	expiryTime := time.Hour * 24 * 7
	token, _ := tToken.CreateToken(data, expiryTime)
	tToken.RefreshToken = SingleToken{
		BearerKey:  token,
		ExpiryTime: expiryTime,
	}
}

// func VerifyToken()
