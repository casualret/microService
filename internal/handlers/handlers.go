package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"microService/internal/service"
)

const (
	incorrectData       = "Некорректные данные"
	userUnauthorized    = "Пользователь не авторизован"
	userAccessDenied    = "Пользователь не имеет доступа"
	bannerNotFound      = "Баннер не найден"
	internalServerError = "Внутренняя ошибка сервера"
	bannerNotSelected   = "Баннер не выбран"
	errorAuthorize      = "Ошибка авторизации"
	errorParseToken     = "Ошибка парса токена"
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
	r := gin.New()
	//r.Use(func(c *gin.Context) {
	//	c.Set("appCtx", h.App)
	//	c.Next()
	//})

	r.POST("/user", h.SignUp)
	r.GET("/user", h.SignIn)

	//r.POST("/tag", h.CreateTag)
	//r.POST("/feature", h.CreateFeature)
	r.Use(h.JWTAuth)
	r.GET("/user_banner", h.GetUserBanner)

	banner := r.Group("/banner")
	banner.Use(h.isAdminMiddleware)
	{
		banner.GET("", h.isAdminMiddleware, h.GetBanners)
		banner.POST("", h.CreateBanner)
		banner.DELETE("/:id", h.DeleteBanner)
		banner.PATCH("/:id", h.ChangeBanner)
	}

	return r
}
