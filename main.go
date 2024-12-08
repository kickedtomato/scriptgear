package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeHandler(c *gin.Context, h gin.H) {

}

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.LoadHTMLGlob("pages/*.html")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/home")
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base", gin.H{
			"HostURL": "",
			"SubPage": "Home",
		})
	})

	r.Run(`:8080`)
}
