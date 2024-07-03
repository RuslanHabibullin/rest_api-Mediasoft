package handler

import (
	"tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.furniture_POST)
			lists.GET("/", h.furnitures_GET)
			lists.GET("/:id", h.furniture_GET)
			lists.PUT("/:id", h.furniture_PUT)
			lists.DELETE("/:id", h.furniture_DELETE)
			lists.PATCH("/:id", h.furniture_PATCH)
		}
	}

	return router
}
