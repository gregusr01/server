package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/server/tokenmanager"
	"github.com/go-vela/server/util"
)

type AuthToken struct {
	JWTToken string
}

func MustValidToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.Request.Header.Get("Authorization")
		claims, err := tokenmanager.FromContext(c).ValidateToken(c, t)
		if err != nil {
			retErr := fmt.Errorf("unable to validate token for refresh: %s", err)
			util.HandleError(c, http.StatusUnauthorized, retErr)
			return
		}
		ToContext(c, claims)
		c.Next()
	}
}

func MustRegistration() gin.HandlerFunc {
	return func(c *gin.Context) {
		cl := FromContext(c)
		if cl.TokenType != "Registration" {
			retErr := fmt.Errorf("the type of token is not of type registration")
			util.HandleError(c, http.StatusUnauthorized, retErr)
			return
		}
		c.Next()
	}
}

func MustAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cl := FromContext(c)
		if cl.TokenType != "Auth" {
			retErr := fmt.Errorf("the type of token is not of type auth")
			util.HandleError(c, http.StatusUnauthorized, retErr)
			return
		}
		c.Next()
	}
}
