package main

import (
	"net/http"

	"mango"
)

func main() {
	r := mango.New()
	r.GET("/", func(c *mango.Context) {
		c.HTML(http.StatusOK, "<h1>Hello mango</h1>")
	})
	r.GET("/hello", func(c *mango.Context) {
		// expect /hello?name=mangouser
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *mango.Context) {
		c.JSON(http.StatusOK, mango.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
