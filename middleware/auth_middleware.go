package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "log in inorder to access this route"})
			ctx.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid token")
			}
			return os.Getenv("JWT_SECRET"), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			ctx.Abort()
			return
		}

		userID := claims["user_id"]
		role := claims["role"]

		ctx.Set("user_id", userID)
		ctx.Set("role", role)

		ctx.Next()
	}
}

func IsAdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, ok := ctx.Get("role")
		if !ok || role != "admin" {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "unauthorized to access this route"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func IsOwnerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		userID, ok := ctx.Get("user_id")

		if !ok || userID != id {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "unauthorized to access this route"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}