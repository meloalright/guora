package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/conf"
	"github.com/meloalright/guora/internal/h"
	"github.com/meloalright/guora/internal/model"
	"github.com/meloalright/guora/internal/service/rdbservice"
	// "strconv"
)

// Index func
func Index(c *gin.Context) {

	var a model.Answer
	var q model.Question
	var hotAnswers []model.Answer
	var hotQuestions []model.Question
	var err error

	if hotAnswers, err = a.GetOrderList(10, 0, "id desc"); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

		return
	}

	PID, exist := c.Get("pid")
	if !exist {
		c.JSON(200, h.Response{Status: 404, Message: "Not exist"})
		c.Abort()
		return
	}

	ProfileID, ok := PID.(int)
	if !ok {
		c.JSON(200, h.Response{Status: 404, Message: "Not int"})
		c.Abort()
		return
	}

	if err = rdbservice.RedisWrapListSupported(hotAnswers, ProfileID); err != nil {

		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return

	}

	if hotQuestions, err = q.GetOrderList(10, 0, "id desc"); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

		return
	}

	csrdata := map[string]interface{}{
		"hotAnswers":       hotAnswers,
		"hotAnswersCounts": len(hotAnswers),
		"hotQuestions":     hotQuestions,
	}

	template := "index.html"
	data := map[string]interface{}{
		"lang":    conf.Config().Lang,
		"csrdata": csrdata,
	}

	c.HTML(http.StatusOK, template, data)

	return
}
