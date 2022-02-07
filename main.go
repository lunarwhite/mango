package main

import (
	"net/http"

	"mango"
)

func main() {
	r := mango.Default()
	r.GET("/", func(c *mango.Context) {
		c.String(http.StatusOK, "Hello mangouser\n")
	})

	// index out of range for testing Recovery()
	r.GET("/panic", func(c *mango.Context) {
		names := []string{"mangouser"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
