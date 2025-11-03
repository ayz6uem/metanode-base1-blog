package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware is a Gin middleware that logs HTTP requests in JSON format using logrus.
func LoggerMiddleware() gin.HandlerFunc {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		log.WithFields(logrus.Fields{
			"client_ip":     param.ClientIP,
			"timestamp":     param.TimeStamp.Format(time.RFC3339),
			"method":        param.Method,
			"path":          param.Path,
			"status_code":   param.StatusCode,
			"latency":       param.Latency,
			"user_agent":    param.Request.UserAgent(),
			"error_message": param.ErrorMessage,
		}).Info("HTTP Request")
		return ""
	})
}

// ErrorHandlingMiddleware is a Gin middleware that recovers from panics and logs the error using logrus.
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logrus.WithFields(logrus.Fields{
					"error":  r,
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
				}).Error("Recovered from panic")

				c.JSON(500, gin.H{
					"code":    500,
					"message": "Recovered from panic",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
