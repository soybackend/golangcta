package main

import (
	"ginhttpdemo/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	sellerRepository := controllers.New()
	r.POST("/seller", sellerRepository.CreateSeller)
	r.GET("/seller", sellerRepository.GetSellers)
	r.GET("/seller/:id", sellerRepository.GetSeller)
	r.PUT("/seller/:id", sellerRepository.UpdateSeller)
	r.DELETE("/seller/:id", sellerRepository.DeleteSeller)

	return r
}