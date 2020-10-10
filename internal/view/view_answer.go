package view

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/conf"
	"github.com/meloalright/guora/internal/h"
	"github.com/meloalright/guora/internal/model"
	"github.com/meloalright/guora/internal/service/rdbservice"
)

// Answer func
func Answer(c *gin.Context) {
	var a model.Answer
	var q model.Question
	var answer model.Answer
	var hotQuestions []model.Question
	var err error

	answerID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	a.ID = answerID

	if answer, err = a.Get(); err != nil {

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

	if err = rdbservice.RedisWrapSupported(&answer, ProfileID); err != nil {

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
		"answer":       answer,
		"hotQuestions": hotQuestions,
	}

	template := "answer.html"
	data := map[string]interface{}{
		"lang":    conf.Config().Lang,
		"csrdata": csrdata,
	}

	c.HTML(http.StatusOK, template, data)

	return
}
