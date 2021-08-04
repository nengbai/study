package prometheus

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
)

func test() {
	router1 := gin.New()
	router1.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "it works")
	})
	p := ginprometheus.NewPrometheus("test")
	router1.Use(p.HandlerFunc())
	router2 := gin.New()
	p.SetMetricsPath(router2)
	go func() { router2.Run(":7100") }()
	router1.Run(":8000")
}
