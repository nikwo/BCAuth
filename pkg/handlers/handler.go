package handlers

import (
	"BCAuth/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services services.Service
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, PATCH, OPTIONS, GET, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	engine := gin.New()
	engine.Use(CORSMiddleware())

	api := engine.Group("/api/v1")
	{
		api.POST("/sign-up", h.Register)
		api.POST("/sign-in", h.SignIn)
		users := api.Group("/users", h.ValidateToken)
		{
			users.GET("/:id", h.UserInfo)
			users.GET("/", h.BrowseUsers)
			users.PATCH("/:id", h.UpdateUser)
			users.DELETE("/:id", h.DeleteUser)
		}
	}

	return engine
}

func HandlerInit(service *services.Service) *Handler {
	return &Handler{
		services: *service,
	}
}

func (h *Handler) Greeting(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "Hello",
	})
}
