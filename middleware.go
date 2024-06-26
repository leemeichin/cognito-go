package cognito

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (cog *Cognito) Authorize(c *gin.Context) {
	tokenHeader, err := tokenFromAuthHeader(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "invalid Authorization header"})
		return
	}
	token, err := cog.VerifyToken(tokenHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	c.Set("Auth", map[string]interface{}{
		"token":         token,
		"email":         token.Claims.(jwt.MapClaims)["email"],
		"emailVerified": token.Claims.(jwt.MapClaims)["email_verified"],
		"sub":           token.Claims.(jwt.MapClaims)["sub"],
	})

	c.Next()
}

func tokenFromAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no token")
	}

	parts := strings.Fields(authHeader)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid Authorization header format")
	}

	return parts[1], nil
}
