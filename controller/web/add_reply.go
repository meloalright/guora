package web

// import "fmt"

import (
	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/h"
	"github.com/meloalright/guora/model"
)

type CreateReplyInterface struct {
	Content          string `json:"content"`
	Type             int    `json:"type"`
	CommentID        int    `json:"commentID"`
	ReplyToProfileID int    `json:"replyToProfileID"`
}

func CreateReply(c *gin.Context) {
	var m CreateReplyInterface
	var r model.Reply
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
	r.ReplyFromProfileID = value

	if err = c.ShouldBindJSON(&m); err != nil {
		c.JSON(200, h.Response{Status: 500, Message: err.Error()})
		return
	}
	r.Content = m.Content
	r.Type = m.Type
	r.CommentID = m.CommentID
	r.ReplyToProfileID = m.ReplyToProfileID
	if ra, err = r.Create(); err != nil {

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
