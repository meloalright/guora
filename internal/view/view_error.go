package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/conf"
)

func ViewError(c *gin.Context) {

	var csrdata map[string]interface{}

	template := "error.html"
	data := map[string]interface{}{
		"lang":    conf.Config().Lang,
		"csrdata": csrdata,
	}

	c.HTML(http.StatusOK, template, data)

	return
}
