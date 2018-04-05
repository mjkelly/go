package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(gin.BasicAuth(gin.Accounts{
		"duplex": "agreements",
	}))

	r.StaticFS("/", http.Dir("."))

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
