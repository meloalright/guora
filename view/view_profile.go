package view

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/h"
	"github.com/meloalright/guora/model"
	"github.com/meloalright/guora/service/rdbservice"
)

func ViewProfile(c *gin.Context) {

	var p model.Profile
	var a model.Answer
	var q model.Question
	var _q model.Question
	var profile model.Profile
	var answers []model.Answer
	var answersCounts int
	var questions []model.Question
	var questionsCounts int
	var hotQuestions []model.Question
	var csrdata map[string]interface{}
	var err error

	profileID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	p.ID = profileID
	a.AnswerProfileID = profileID
	q.QuestionProfileID = profileID

	if profile, err = p.Get(); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})
		return

	}

	if hotQuestions, err = _q.GetOrderList(10, 0, "id desc"); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})
		return

	}

	if answersCounts, err = a.GetCounts(); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})
		return

	}

	if questionsCounts, err = q.GetCounts(); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})
		return

	}

	card := c.Query("card")
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	switch card {
	case "question":

		if questions, err = q.GetList(10, 0); err != nil {

			c.JSON(200, h.Response{
				Status:  404,
				Message: err.Error(),
			})
			return

		}

		csrdata = map[string]interface{}{
			"profile":         profile,
			"questions":       questions,
			"answersCounts":   answersCounts,
			"questionsCounts": questionsCounts,
			"hotQuestions":    hotQuestions,
		}

	case "document":

		csrdata = map[string]interface{}{
			"profile":         profile,
			"answersCounts":   answersCounts,
			"questionsCounts": questionsCounts,
			"hotQuestions":    hotQuestions,
		}

	default:

		if answers, err = a.GetList(10, 0); err != nil {

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

		if err = rdbservice.RedisWrapListSupported(answers, ProfileID); err != nil {

			c.JSON(200, h.Response{
				Status:  500,
				Message: err.Error(),
			})
			return

		}

		csrdata = map[string]interface{}{
			"profile":         profile,
			"answers":         answers,
			"answersCounts":   answersCounts,
			"questionsCounts": questionsCounts,
			"hotQuestions":    hotQuestions,
		}

	}

	template := "profile.html"
	data := map[string]interface{}{
		"csrdata": csrdata,
	}

	c.HTML(http.StatusOK, template, data)

	return
}
