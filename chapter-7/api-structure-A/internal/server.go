package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Timing-Allow-Origin", "*")
		// Handle browser preflight requests, where it asks for allowed origin.
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func WithCustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				_ = c.Error(fmt.Errorf("[recovered] panic: %v", err))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
			}
		}()
		c.Next()
	}
}

func BuildEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(CORS())
	group := engine.Group("/api")
	RegisterRoutes(group)
	// engine.use(<any middleware you want>)
	return engine
}
