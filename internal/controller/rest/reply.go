package rest

// import "fmt"

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/internal/h"
	"github.com/meloalright/guora/internal/model"
)

// GetReply func
func GetReply(c *gin.Context) {

	var r model.Reply
	var err error

	var id = c.Param("id")

	r.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if reply, err := r.Get(); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

	} else {

		c.JSON(200, h.Response{
			Status:  200,
			Message: reply,
		})

	}

	return
}

// UpdateReply func
func UpdateReply(c *gin.Context) {

	var r model.Reply
	var ra int64
	var err error

	var id = c.Param("id")

	r.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&r); err != nil {
		c.JSON(200, h.Response{Status: 500, Message: err.Error()})
		return
	}

	if ra, err = r.Update(); err != nil {

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

// DeleteReply func
func DeleteReply(c *gin.Context) {

	var r model.Reply
	var ra int64
	var err error

	var id = c.Param("id")

	r.ID, err = strconv.Atoi(id)

	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if ra, err = r.Delete(); err != nil {

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

// GetReplies func
func GetReplies(c *gin.Context) {

	var r model.Reply
	var limit int
	var offset int
	var err error

	commentID, _ := strconv.Atoi(c.Query("commentID"))
	r.CommentID = commentID

	limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}
	offset, err = strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if replies, err := r.GetList(limit, offset); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

	} else {

		c.JSON(200, h.Response{
			Status:  200,
			Message: replies,
		})

	}

	return
}

// GetRepliesCounts func
func GetRepliesCounts(c *gin.Context) {

	var r model.Reply

	commentID, _ := strconv.Atoi(c.Query("commentID"))
	r.CommentID = commentID

	if counts, err := r.GetCounts(); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

	} else {

		c.JSON(200, h.Response{
			Status:  200,
			Message: counts,
		})

	}

	return
}
