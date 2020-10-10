package rest

// import "fmt"

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/internal/h"
	"github.com/meloalright/guora/internal/model"
)

// GetProfile func
func GetProfile(c *gin.Context) {

	var p model.Profile
	var err error

	var id = c.Param("id")

	p.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if profile, err := p.Get(); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

	} else {

		c.JSON(200, h.Response{
			Status:  200,
			Message: profile,
		})

	}

	return
}

// UpdateProfile func
func UpdateProfile(c *gin.Context) {

	var p model.Profile
	var ra int64
	var err error

	var id = c.Param("id")

	p.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&p); err != nil {
		c.JSON(200, h.Response{Status: 500, Message: err.Error()})
		return
	}
	if ra, err = p.Update(); err != nil {

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

// DeleteProfile func
func DeleteProfile(c *gin.Context) {

	var p model.Profile
	var ra int64
	var err error

	var id = c.Param("id")

	p.ID, err = strconv.Atoi(id)

	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if ra, err = p.Delete(); err != nil {

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

// GetProfiles func
func GetProfiles(c *gin.Context) {

	var p model.Profile
	var limit int
	var offset int
	var err error

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

	if profiles, err := p.GetList(limit, offset); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

	} else {

		c.JSON(200, h.Response{
			Status:  200,
			Message: profiles,
		})

	}

	return
}

// GetProfilesCounts func
func GetProfilesCounts(c *gin.Context) {

	var p model.Profile

	if counts, err := p.GetCounts(); err != nil {

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
