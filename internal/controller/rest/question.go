package rest

// import "fmt"

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/internal/h"
	"github.com/meloalright/guora/internal/model"
)

// GetQuestion func
func GetQuestion(c *gin.Context) {

	var q model.Question
	var err error

	var id = c.Param("id")

	q.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if question, err := q.Get(); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

	} else {

		c.JSON(200, h.Response{
			Status:  200,
			Message: question,
		})

	}

	return
}

// UpdateQuestion func
func UpdateQuestion(c *gin.Context) {

	var q model.Question
	var ra int64
	var err error

	var id = c.Param("id")

	q.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&q); err != nil {
		c.JSON(200, h.Response{Status: 500, Message: err.Error()})
		return
	}

	if ra, err = q.Update(); err != nil {

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

// DeleteQuestion func
func DeleteQuestion(c *gin.Context) {

	var q model.Question
	var ra int64
	var err error

	var id = c.Param("id")

	q.ID, err = strconv.Atoi(id)

	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if ra, err = q.Delete(); err != nil {

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

// GetQuestions func
func GetQuestions(c *gin.Context) {

	var q model.Question
	var limit int
	var offset int
	var err error

	questionProfileID, _ := strconv.Atoi(c.Query("questionProfileID"))
	q.QuestionProfileID = questionProfileID

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

	if questions, err := q.GetList(limit, offset); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

	} else {

		c.JSON(200, h.Response{
			Status:  200,
			Message: questions,
		})

	}

	return
}

// GetQuestionsCounts func
func GetQuestionsCounts(c *gin.Context) {

	var q model.Question

	questionProfileID, _ := strconv.Atoi(c.Query("questionProfileID"))
	q.QuestionProfileID = questionProfileID

	if counts, err := q.GetCounts(); err != nil {

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
