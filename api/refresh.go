package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/server/router/middleware/auth"
	"github.com/go-vela/server/tokenmanager"
	"github.com/go-vela/server/util"
)

type AuthToken struct {
	JWTToken string
}

func SystemRefresh(c *gin.Context) {
	cl := auth.FromContext(c)
	t := c.Request.Header.Get("Authorization")
	nt, err := tokenmanager.FromContext(c).MintToken(c, "Auth", *cl.Sub)
	if err != nil {
		retErr := fmt.Errorf("unable to mint new token for refresh: %s", err)
		util.HandleError(c, http.StatusUnauthorized, retErr)
		return
	}
	at := AuthToken{
		JWTToken: nt,
	}
	if err = tokenmanager.FromContext(c).InvalidateToken(c, t); err != nil {
		retErr := fmt.Errorf("unable add token for to invalidation db: %s", err)
		util.HandleError(c, http.StatusUnauthorized, retErr)
		return
	}
	c.JSON(http.StatusOK, at)
}
