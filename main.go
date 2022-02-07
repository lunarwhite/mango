package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"mango"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := mango.New()
	r.Use(mango.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "User", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}

	r.GET("/", func(c *mango.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *mango.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", mango.H{
			"title":  "mango",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.Run(":9999")
}
