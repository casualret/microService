package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	//config cleanenv
	//db	pq
}

func main() {

	r := InitRoutes()
	err := r.Run("localhost:8084")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Пока я мастерил фрегат мир стал бессмыслено богат и полон гнуси")
	})
	//r.POST("/create_url", CreateShortUrl(c))
	//r.GET("/get_url", RedirectUrl(c))
	return r
}
