package infrastructure

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		token, err := ValidateJwtToken(authHeader)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
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