package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"urlshortener/internal/service"
)

type Handlers struct {
	App    *service.App
	Logger *slog.Logger
}

func NewHandlers(service *service.App, logger *slog.Logger) *Handlers {
	return &Handlers{
		App:    service,
		Logger: logger,
	}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	//r := gin.Default()
	r := gin.New()

	r.GET("/user_banner", func(c *gin.Context) {
		c.String(200, "Пока я мастерил фрегат мир стал бессмыслено богат и полон гнуси")
	})
	r.POST("/banner", h.CreateBanner)

	return r
}
