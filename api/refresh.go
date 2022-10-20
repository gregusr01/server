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
	claims, err := tokenmanager.ValidateToken(t)
	if err != nil {
		retErr := fmt.Errorf("unable to validate token for refresh: %s", err)
		util.HandleError(c, http.StatusUnauthorized, retErr)
		return
	}
	nt, err := tokenmanager.MintToken("Auth", claims.Sub)
	if err != nil {
		retErr := fmt.Errorf("unable to mint new token for refresh: %s", err)
		util.HandleError(c, http.StatusUnauthorized, retErr)
		return
	}
	at := AuthToken{
		JWTToken: nt,
	}
	c.JSON(http.StatusOK, at)
}
