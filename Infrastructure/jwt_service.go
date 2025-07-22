package infrastructure

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/abeni-al7/task_manager/Domain"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJwtToken(user *domain.User, password string) (string, error) {
	err := ComparePassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email": user.Email,
		"username": user.Username,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", errors.New("unable to generate token")
	}

	return jwtToken, nil
}

func ValidateJwtToken(authHeader string) (*jwt.Token, error) {
	if authHeader == "" {
		return &jwt.Token{}, errors.New("log in inorder to access this route")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return &jwt.Token{}, errors.New("invalid authorization header")
	}

	token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return &jwt.Token{}, errors.New("invalid token")
	}
	return token, nil
}