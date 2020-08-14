package web

import (
	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/h"
	"github.com/meloalright/guora/model"
)

type CreateQuestionInterface struct {
	Title string `json:"title"`
	Type  int    `json:"type"`
	Desc  string `json:"desc"`
}

func CreateQuestion(c *gin.Context) {
	var m CreateQuestionInterface
	var q model.Question
	var ra int64
	var err error

	ProfileID, exist := c.Get("pid")
	if !exist {
		c.JSON(200, h.Response{Status: 404, Message: "Not exist"})
		c.Abort()
		return
	}
	value, ok := ProfileID.(int)
	if !ok {
		c.JSON(200, h.Response{Status: 404, Message: "Not int"})
		c.Abort()
		return
	}
	q.QuestionProfileID = value

	if err = c.ShouldBindJSON(&m); err != nil {
		c.JSON(200, h.Response{Status: 500, Message: err.Error()})
		return
	}
	q.Title = m.Title
	q.Type = m.Type
	q.Desc = m.Desc
	ra, err = q.Create()

	if err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

	} else {

		c.JSON(200, h.Response{
			Status: 200,
			Message: map[string]interface{}{
				"record": q,
				"ra":     ra,
			},
		})

	}

	return
}
