package tinytoken

import (
	"errors"
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
	SecretKey    string
}

func (tToken *TinyToken) Genera(data interface{}) {

	tToken.AccessTokenGenerator(data)
	tToken.RefreshTokenGenerator(data)
}

func (tToken *TinyToken) CreateToken(data interface{}, expiryTime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(expiryTime).Unix(),
	})

	var secretKey = []byte(tToken.SecretKey)
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

func VerifyToken(tokenString string, secretKey string) (jwt.MapClaims, error) {

	var secretKeyInByte = []byte(secretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKeyInByte, nil
	})

	if err != nil {
		return nil, err // Token parsing error
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return nil, errors.New("token is expired")
			}
		}
		return claims, nil // Token is valid
	}
	return nil, errors.New("invalid token")
}
