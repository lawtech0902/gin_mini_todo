package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/gin_todo_demo/backend/pkg/app"
	"github.com/lawtech0902/gin_todo_demo/backend/pkg/e"
	"net/http"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// parse the jwt token
		if err := app.ParseRequest(c); err != nil {
			app.SendResponse(c, http.StatusBadRequest, e.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		
		c.Next()
	}
}

func NoCacheMiddleware(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

func CorsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	
	c.Next()
}

func SecureMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
}
