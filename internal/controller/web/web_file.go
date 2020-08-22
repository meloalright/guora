package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/internal/constant"
	"github.com/meloalright/guora/internal/h"
)

var AvatarDir = "./web/static/avatar"

func FileAvatarResolve(c *gin.Context) {

	ProfileID, exist := c.Get("pid")
	if !exist {
		c.JSON(200, h.Response{Status: 404, Message: "Not exist"})
		return
	}

	value, ok := ProfileID.(int)
	if !ok {
		c.JSON(200, h.Response{Status: 404, Message: "Not int"})
		c.Abort()
		return
	}

	file, err := c.FormFile(constant.AVATARKEY)
	if err != nil {

		c.JSON(200, h.Response{
			Status:  500,
			Message: err,
		})
		return
	}

	fileReference := fmt.Sprintf("%v%v%v%v", AvatarDir, "/", value, ".png")

	if err := c.SaveUploadedFile(file, fileReference); err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err,
		})
	} else {
		c.JSON(200, h.Response{
			Status:  200,
			Message: fileReference,
		})
	}

	return
}
