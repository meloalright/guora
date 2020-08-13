package web

// import "fmt"

import (
	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/h"
	"github.com/meloalright/guora/model"
)

type CreateCommentInterface struct {
	Content  string `json:"content"`
	Type     int    `json:"type"`
	AnswerID int    `json:"answerID"`
}

func CreateComment(c *gin.Context) {
	var m CreateCommentInterface
	var co model.Comment
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

	co.CommentProfileID = value

	if err = c.ShouldBindJSON(&m); err != nil {
		c.JSON(200, h.Response{Status: 500, Message: err.Error()})
		return
	}
	co.Content = m.Content
	co.Type = m.Type
	co.AnswerID = m.AnswerID

	if ra, err = co.Create(); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

	} else {

		c.JSON(200, h.Response{
			Status:  200,
			Message: ra,
		})

	}

	return
}
