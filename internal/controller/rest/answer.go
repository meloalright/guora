package rest

// import "fmt"

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/internal/h"
	"github.com/meloalright/guora/internal/model"
	"github.com/meloalright/guora/internal/service/rdbservice"
)

var ctx = context.Background()

func GetAnswer(c *gin.Context) {

	var a model.Answer
	var answer model.Answer
	var err error

	var id = c.Param("id")

	a.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

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

	} else {
		c.JSON(200, h.Response{
			Status:  200,
			Message: answer,
		})
	}

	return
}

func UpdateAnswer(c *gin.Context) {

	var a model.Answer
	var ra int64
	var err error

	var id = c.Param("id")

	a.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&a); err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if ra, err = a.Update(); err != nil {

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

func DeleteAnswer(c *gin.Context) {

	var a model.Answer
	var ra int64
	var err error

	var id = c.Param("id")

	a.ID, err = strconv.Atoi(id)

	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if ra, err = a.Delete(); err != nil {

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

func GetAnswers(c *gin.Context) {

	var a model.Answer
	var limit int
	var offset int
	var err error
	var answers []model.Answer

	questionID, _ := strconv.Atoi(c.Query("questionID"))
	a.QuestionID = questionID

	answerProfileID, _ := strconv.Atoi(c.Query("answerProfileID"))
	a.AnswerProfileID = answerProfileID

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

	if c.Query("order") == "supporters_counts" {

		if answers, err = a.GetOrderList(limit, offset, "supporters_counts desc"); err != nil {

			c.JSON(200, h.Response{
				Status:  500,
				Message: err.Error(),
			})

			return
		}

	} else {

		if answers, err = a.GetList(limit, offset); err != nil {

			c.JSON(200, h.Response{
				Status:  500,
				Message: err.Error(),
			})

			return
		}
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

	if err = rdbservice.RedisWrapListSupported(answers, ProfileID); err != nil {

		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})

	} else {

		c.JSON(200, h.Response{
			Status:  200,
			Message: answers,
		})

	}

	return
}

func GetAnswersCounts(c *gin.Context) {

	var a model.Answer

	questionID, _ := strconv.Atoi(c.Query("questionID"))
	a.QuestionID = questionID

	answerProfileID, _ := strconv.Atoi(c.Query("answerProfileID"))
	a.AnswerProfileID = answerProfileID

	if counts, err := a.GetCounts(); err != nil {

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
