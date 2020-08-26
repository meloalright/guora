package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/conf"
)

func ViewAdmin(c *gin.Context) {

	var csrdata map[string]interface{}

	template := "admin.html"
	data := map[string]interface{}{
		"lang":    conf.Config().Lang,
		"csrdata": csrdata,
	}

	c.HTML(http.StatusOK, template, data)

	return
}
