package api

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

func SystemRefresh(c *gin.Context) {
	t := c.Request.Header.Get("Authorization")
	claims, err := tokenmanager.FromContext(c).ValidateToken(c, t)
	if err != nil {
		retErr := fmt.Errorf("unable to validate token for refresh: %s", err)
		util.HandleError(c, http.StatusUnauthorized, retErr)
		return
	}

	//if claims.TokenType == "Registration"
	// then register worker

	nt, err := tokenmanager.FromContext(c).MintToken(c, "Auth", claims.Sub)
	if err != nil {
		retErr := fmt.Errorf("unable to mint new token for refresh: %s", err)
		util.HandleError(c, http.StatusUnauthorized, retErr)
		return
	}
	at := AuthToken{
		JWTToken: nt,
	}

	// invalidate token
	if err = tokenmanager.FromContext(c).InvalidateToken(c, t); err != nil {
		retErr := fmt.Errorf("unable add token for to invalidation db: %s", err)
		util.HandleError(c, http.StatusUnauthorized, retErr)
		return
	}

	c.JSON(http.StatusOK, at)
}
