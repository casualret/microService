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

//func JWTAuth() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		const BearerSchema = "Bearer "
//		header := c.GetHeader("Authorization")
//		if header == "" {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
//			c.Abort()
//			return
//		}
//
//		if !strings.HasPrefix(header, BearerSchema) {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Header"})
//			c.Abort()
//			return
//		}
//
//		tokenStr := header[len(BearerSchema):]
//		claims := &auth.Claims{}
//
//		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
//			return auth.JwtKey, nil
//		})
//
//		if err != nil {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
//			c.Abort()
//			return
//		}
//
//		if !token.Valid {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
//			c.Abort()
//			return
//		}
//
//		c.Set("username", claims.Username)
//		c.Next()
//	}
//}

func (h *Handlers) InitRoutes() *gin.Engine {
	r := gin.New()
	//r.Use(func(c *gin.Context) {
	//	c.Set("appCtx", h.App)
	//	c.Next()
	//})

	r.POST("/user", h.SignUp)
	r.GET("/user", h.SignIn)

	r.POST("/tag", h.CreateTag)
	r.POST("/feature", h.CreateFeature)

	banner := r.Group("/banner")
	banner.Use(h.JWTAuth)
	{
		banner.GET("", h.GetBanners)
		banner.POST("", h.CreateBanner)
		banner.DELETE("/:id", h.DeleteBanner)
		banner.PATCH("/:id", h.ChangeBanner)
	}

	r.GET("/user_banner", h.GetUserBanner)

	return r
}
