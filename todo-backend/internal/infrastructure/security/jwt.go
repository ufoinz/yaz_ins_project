package security

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// generateToken creates a JWT with a "user_id" field and a 24-hour lifetime
func GenerateToken(userID int64, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// JWTMiddleware returns Gin middleware
func JWTMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extracting the value of the "Authorization" header
		header := c.GetHeader("Authorization")
		parts := strings.SplitN(header, " ", 2)

		// expecting the "Bearer <token>" format
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}

		// Parse and verify the signed token
		tok, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(secret), nil
		})
		if err != nil || !tok.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// Extract a set of fields (claims) and set the user_id to int64
		claims := tok.Claims.(jwt.MapClaims)
		id := int64(claims["user_id"].(float64))

		// Saving the user_id in the context for subsequent handlers
		c.Set("user_id", id)

		// Continue the chain execution
		c.Next()
	}
}
