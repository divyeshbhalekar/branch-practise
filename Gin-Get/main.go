package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func postpage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func querystring(c *gin.Context) {
	name := c.Query("name")
	email := c.Query("email")

	c.JSON(200, gin.H{
		"name":  name,
		"email": email,
	})
}
func pathparameter(c *gin.Context) {
	name := c.Param("name")
	email := c.Param("email")

	c.JSON(200, gin.H{
		"name":  name,
		"email": email,
	})
}

func main() {

	r := gin.Default()
	r.GET("/", Homepage)
	r.POST("/post", postpage)
	r.GET("/query", querystring)
	r.GET("/path/:name/:age", pathparameter)

	r.Run(":3001")
}
