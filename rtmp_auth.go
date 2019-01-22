package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	runGin()
}

func runGin() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/publish", Publish)
		v1.POST("/unpublish", UnPublish)
		v1.POST("/connect", Connect)
		v1.POST("/close", Close)
		v1.GET("/smoke", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})
	}

	if err := router.Run(":1188"); err != nil {
		log.Panic("Listen and serve at 1188 failed")
	}
}
