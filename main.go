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
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *mango.Context) {
		// expect /hello/mangouser
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *mango.Context) {
		c.JSON(http.StatusOK, mango.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
