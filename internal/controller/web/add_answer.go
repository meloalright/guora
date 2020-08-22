package web

import (
	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/internal/h"
	"github.com/meloalright/guora/internal/model"
)

type CreateAnswerInterface struct {
	Content    string `json:"content"`
	Type       int    `json:"type"`
	QuestionID int    `json:"questionID"`
}

func CreateAnswer(c *gin.Context) {
	var m CreateAnswerInterface
	var a model.Answer
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

	a.AnswerProfileID = value

	if err = c.ShouldBindJSON(&m); err != nil {
		c.JSON(200, h.Response{Status: 500, Message: err.Error()})
		return
	}

	a.Content = m.Content
	a.Type = m.Type
	a.QuestionID = m.QuestionID

	if ra, err = a.Create(); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

	} else {

		c.JSON(200, h.Response{
			Status: 200,
			Message: map[string]interface{}{
				"record": a,
				"ra":     ra,
			},
		})

	}

	return
}
