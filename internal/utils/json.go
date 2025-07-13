package utils

import (
	"test-case-vhiweb/internal/dtos"

	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, status int,
	data any) {
	r := &dtos.ResponseDTO{
		Success: true,
		Data:    data,
	}

	c.JSON(status, r)
}

func JSONError(c *gin.Context, status int,
	err any) {
	r := &dtos.ResponseDTO{
		Success: false,
		Error:   err,
	}

	c.AbortWithStatusJSON(status, r)
}
