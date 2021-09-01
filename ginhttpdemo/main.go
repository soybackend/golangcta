package main

import (
	"fmt"
	"ginhttpdemo/controllers"
	"github.com/gin-gonic/gin"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
	"os"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("CTA App"),
		newrelic.ConfigLicense("ef5a0bef83707727aab0557e9622d738e35aNRAL"),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	r.Use(nrgin.Middleware(app))

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