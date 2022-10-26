package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/server/router/middleware/org"
	"github.com/go-vela/server/router/middleware/user"
	"github.com/go-vela/server/scm"
	"github.com/go-vela/server/tokenmanager"
	"github.com/go-vela/server/util"
	"github.com/sirupsen/logrus"
)

type AuthToken struct {
	JWTToken string
}

func Register(c *gin.Context) {

	// validate User via git token
	// return registration token

	o := org.Retrieve(c)
	u := user.Retrieve(c)

	c.Request.ParseForm()

	hn := c.Request.Form.Get("hostname")

	perm, err := scm.FromContext(c).OrgAccess(u, o)
	if err != nil {
		logrus.Errorf("unable to get user %s access level for org %s", u.GetName(), o)
	}

	if perm != "admin" {
		//reject
	}

	ntk, err := tokenmanager.FromContext(c).MintToken(c, "Registration", hn)
	if err != nil {
		retErr := fmt.Errorf("unable to validate token for registration: %s", err)
		util.HandleError(c, http.StatusUnauthorized, retErr)
		return
	}

	at := AuthToken{
		JWTToken: ntk,
	}

	c.JSON(http.StatusOK, at)
}
