package main

import (
	"net/http"

	"mango"
)

func main() {
	r := mango.New()
	r.GET("/index", func(c *mango.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *mango.Context) {
			c.HTML(http.StatusOK, "<h1>Hello mango</h1>")
		})

		v1.GET("/hello", func(c *mango.Context) {
			// expect /hello?name=mangouser
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *mango.Context) {
			// expect /hello/mangouser
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *mango.Context) {
			c.JSON(http.StatusOK, mango.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":9999")
}
