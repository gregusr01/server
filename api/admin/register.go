package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/server/tokenmanager"
	"github.com/go-vela/server/util"
)

type Message struct {
	message string
}

func Register(c *gin.Context) {

	tk := c.Request.Header.Get("Authorization")
	claims, err := tokenmanager.ValidateToken(tk)
	if err != nil {
		retErr := fmt.Errorf("unable to validate token for registration: %s", err)
		util.HandleError(c, http.StatusUnauthorized, retErr)
		return
	}

	//Db call to update registration token

	ntk, err := tokenmanager.MintToken("Auth", claims.Sub)
	if err != nil {
		retErr := fmt.Errorf("unable to validate token for registration: %s", err)
		util.HandleError(c, http.StatusUnauthorized, retErr)
		return
	}

	c.JSON(http.StatusOK, Message{message: "worker successfully registered"})
}
