package middleware

// import "fmt"

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/constant"
	"github.com/meloalright/guora/service/authorization"
)

func Authorizer() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get the ss(signed string) inside cookie
		ss, err := c.Cookie(constant.SSKEY)
		if err != nil {
			log.Print("Authorizer error: ", err)
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		// parse a ss(signed string)
		ID, ProfileID, err := authorization.Parse(ss)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		c.Set("uid", ID)
		c.Set("pid", ProfileID)

		// before request

		c.Next()

		// after request

		log.Print("UID: ", ID)
		log.Print("PID: ", ProfileID)
	}
}
