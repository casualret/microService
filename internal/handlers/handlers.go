package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"urlshortener/internal/service"
)

const (
	incorrectData       = "Некорректные данные"
	userUnauthorized    = "Пользователь не авторизован"
	userAccessDenied    = "Пользователь не имеет доступа"
	bannerNotFound      = "Баннер не найден"
	internalServerError = "Внутренняя ошибка сервера"
)

type Handlers struct {
	App    *service.App
	Logger *slog.Logger
}

type Response struct {
	//Message string `json:"message,omitempty"`
	Error string `json:"error,omitempty"`
}

func Error(msg string) Response {
	return Response{
		Error: msg,
	}
}

func newErrorResponse(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, Error(message))
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

	//r.GET("/user_banner", func(c *gin.Context) {
	//	c.String(200, "Пока я мастерил фрегат мир стал бессмыслено богат и полон гнуси")
	//})
	//r.POST("/banner", h.CreateBanner)
	r.POST("/tag", h.CreateTag)
	r.POST("/feature", h.CreateFeature)
	r.POST("/banner", h.CreateBanner)
	r.GET("/banners", h.GetBanners)
	r.GET("/user_banner", h.GetUserBanner)
	r.GET("/test_s", func(c *gin.Context) {
		// Получаем данные из запроса
		steps := c.Query("steps")
		direction := c.Query("direction")

		// Создаем структуру для ответа
		data := gin.H{
			"steps":     steps,
			"direction": direction,
		}

		// Отправляем JSON-ответ
		c.JSON(http.StatusOK, data)
	})

	return r
}
