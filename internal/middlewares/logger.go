package middlewares

import (
	"test-case-vhiweb/internal/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		logger.Log.WithFields(
			logrus.Fields{
				"method":   c.Request.Method,
				"path":     c.Request.URL.Path,
				"duration": time.Since(start),
				"status":   c.Writer.Status(),
			}).Info("Request processed")
	}
}
