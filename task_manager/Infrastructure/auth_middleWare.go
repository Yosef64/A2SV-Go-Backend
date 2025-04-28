package Infrastructure

import (
	"fmt"
	"net/http"
	"strings"
	domain "task_manager/Domain"

	"github.com/gin-gonic/gin"
)

type AuthMiddleWare interface {
	AuthMiddleware() gin.HandlerFunc
	AdminOnly() gin.HandlerFunc
}

type authMiddleWare struct {
	secretKey  string
	jwtService JWTService
}

func NewAuthMiddleware(secretKey string, jwtService JWTService) AuthMiddleWare {
	return &authMiddleWare{
		secretKey:  secretKey,
		jwtService: jwtService,
	}
}
func (a *authMiddleWare) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &domain.Claims{}

		if isValid, err := a.jwtService.ValidateToken(tokenString, claims); err != nil || !isValid {
			if err != nil {
				fmt.Println("Error validating token:", err)
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func (a *authMiddleWare) AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Role not found"})
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok || roleStr != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}
