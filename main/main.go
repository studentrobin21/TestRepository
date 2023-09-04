package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("htmls/*")
	r.Static("/static", "./static")
	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", gin.H{})
	})
	r.GET("/test2", func(c *gin.Context) {
		c.HTML(200, "header.html", gin.H{})
	})
	err := r.Run(":80")
	if err != nil {
		return
	}
}
