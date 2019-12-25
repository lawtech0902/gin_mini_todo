package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/gin_todo_demo/backend/middleware"
	"github.com/lawtech0902/gin_todo_demo/backend/routers/api"
	v1 "github.com/lawtech0902/gin_todo_demo/backend/routers/api/v1"
	"net/http"
)

// InitRouter init router
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CorsMiddleware)
	r.Use(middleware.NoCacheMiddleware)
	r.Use(middleware.SecureMiddleware)
	
	// 404 handler
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	
	r.POST("/token", api.Token)
	
	apiV1 := r.Group("/v1/todos")
	apiV1.Use(middleware.AuthMiddleware())
	{
		apiV1.GET("/", v1.FetchAllTodo)
		apiV1.POST("/", v1.AddTodo)
		apiV1.GET("/:id", v1.FetchOneTodo)
		apiV1.PUT("/:id", v1.UpdateTodo)
		apiV1.DELETE("/:id", v1.DeleteTodo)
	}
	
	svcd := r.Group("/sd")
	
	{
		svcd.GET("/health", v1.HealthCheck)
		svcd.GET("/disk", v1.DiskCheck)
		svcd.GET("/cpu", v1.CPUCheck)
		svcd.GET("/ram", v1.RAMCheck)
	}
	
	return r
}
