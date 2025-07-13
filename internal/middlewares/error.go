package middlewares

import (
	"errors"
	"net/http"
	"test-case-vhiweb/internal/constants"
	"test-case-vhiweb/internal/dtos"
	"test-case-vhiweb/internal/logger"
	"test-case-vhiweb/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors[0]

		logger.Log.Errorf("Error: %s\n", err.Error())

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			validationErrors := make([]dtos.ErrorResponseDTO, 0)
			for _, fe := range ve {
				validationErrors = append(validationErrors, dtos.ErrorResponseDTO{
					Field:   fe.Field(),
					Message: fe.Error(),
				})
			}
			utils.JSONError(c, http.StatusBadRequest, validationErrors)
			return
		}

		var customErr *constants.CustomError
		if errors.As(err, &customErr) {
			httpStatus := constants.ToHTTPStatus(customErr.Code)
			utils.JSONError(c, httpStatus, customErr.Message)
			return
		}

		utils.JSONError(c, http.StatusInternalServerError, constants.Err)
	}
}
