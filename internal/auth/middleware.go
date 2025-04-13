package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RoleModerator = "moderator"
	RoleEmployee  = "employee"
)

func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
			return
		}

		tokenString = tokenString[len("Bearer "):]

		claims, err := ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Невалидный токен: " + err.Error()})
			return
		}

		userRole := claims.Role

		for _, role := range roles {
			if userRole == role {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Доступ запрещен"})
	}
}
