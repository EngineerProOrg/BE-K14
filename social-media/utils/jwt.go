package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "SecretKey_Soci@lMedia"

func GenerateAccessToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"userId":   userId,
		"username": GetUsernameFromEmail(email),
		"exp":      time.Now().Add(time.Hour / 2).Unix(),
	})

	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyToken(accessToken string) (int64, string, error) {
	// if struct implement from interface, we can cast struct
	// token.Method is interface SigningMethod
	// SigningMethodHMAC is a struct that implments interface SigningMethod
	parsedToken, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {

		// check whether your access token was signed with SigningMethodHS256
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return 0, "", errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return 0, "", errors.New("invalid token")
	}

	// data extraction from token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, "", errors.New("invalid token claims")
	}

	// type checking syntax
	userId := (int64)(claims["userId"].(float64))
	username := (string)(claims["username"].(string))

	return userId, username, nil
}
