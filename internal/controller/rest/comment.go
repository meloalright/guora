package rest

// import "fmt"

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/internal/h"
	"github.com/meloalright/guora/internal/model"
)

func GetComment(c *gin.Context) {

	var co model.Comment
	var err error

	var id = c.Param("id")

	co.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if comment, err := co.Get(); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

	} else {

		c.JSON(200, h.Response{
			Status:  200,
			Message: comment,
		})

	}

	return
}

func UpdateComment(c *gin.Context) {

	var co model.Comment
	var ra int64
	var err error

	var id = c.Param("id")

	co.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&co); err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if ra, err = co.Update(); err != nil {

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

func DeleteComment(c *gin.Context) {

	var co model.Comment
	var ra int64
	var err error

	var id = c.Param("id")

	co.ID, err = strconv.Atoi(id)

	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if ra, err = co.Delete(); err != nil {
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

func GetComments(c *gin.Context) {

	var co model.Comment
	var limit int
	var offset int
	var err error

	answerID, _ := strconv.Atoi(c.Query("answerID"))
	co.AnswerID = answerID

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

	if comments, err := co.GetList(limit, offset); err != nil {
		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

	} else {
		c.JSON(200, h.Response{
			Status:  200,
			Message: comments,
		})

	}

	return
}

func GetCommentsCounts(c *gin.Context) {

	var co model.Comment

	answerID, _ := strconv.Atoi(c.Query("answerID"))
	co.AnswerID = answerID

	if counts, err := co.GetCounts(); err != nil {
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
