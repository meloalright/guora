package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/conf"
)

func ViewLogin(c *gin.Context) {

	var csrdata map[string]interface{}

	template := "login.html"
	data := map[string]interface{}{
		"lang":    conf.Config().Lang,
		"csrdata": csrdata,
	}

	c.HTML(http.StatusOK, template, data)

	return
}
