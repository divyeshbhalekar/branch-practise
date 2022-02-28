package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/ping", PingGet)

	router.Run(":3000")
}

func PingGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"Hello": "found Me",
	})
}
