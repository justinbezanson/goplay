package main

import (
	"goplay/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, tmpl templ.Component) error {
	c.Status(status)
	return tmpl.Render(c.Request.Context(), c.Writer)
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		render(c, 200, views.Hello())
	})

	r.Run(":8080")
}
