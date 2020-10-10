package web

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/internal/h"
	"github.com/meloalright/guora/internal/model"
	"github.com/meloalright/guora/internal/service/rdbservice"
)

// CreateSupporter func
func CreateSupporter(c *gin.Context) {
	var su model.Supporter
	var ra int64
	var err error
	var ok bool

	PID, exist := c.Get("pid")
	if !exist {
		c.JSON(200, h.Response{Status: 404, Message: "Not exist"})
		c.Abort()
		return
	}

	su.ProfileID, ok = PID.(int)
	if !ok {
		c.JSON(200, h.Response{Status: 404, Message: "Not int"})
		c.Abort()
		return
	}

	var id = c.Param("id")

	su.AnswerID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if err = rdbservice.RedisAddSupporter(su.AnswerID, su.ProfileID); err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if ra, err = su.Create(); err != nil {

		c.JSON(200, h.Response{
			Status:  500,
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

// DeleteSupporter func
func DeleteSupporter(c *gin.Context) {
	var su model.Supporter
	var ra int64
	var err error
	var ok bool

	PID, exist := c.Get("pid")
	if !exist {
		c.JSON(200, h.Response{Status: 404, Message: "Not exist"})
		c.Abort()
		return
	}

	su.ProfileID, ok = PID.(int)
	if !ok {
		c.JSON(200, h.Response{Status: 404, Message: "Not int"})
		c.Abort()
		return
	}

	var id = c.Param("id")

	su.AnswerID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if err = rdbservice.RedisRemoveSupporter(su.AnswerID, su.ProfileID); err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if ra, err = su.Delete(); err != nil {

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
