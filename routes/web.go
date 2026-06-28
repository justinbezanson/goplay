package routes

import (
	"goplay/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, tmpl templ.Component) error {
	c.Status(status)
	return tmpl.Render(c.Request.Context(), c.Writer)
}

func RegisterWebRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		render(c, 200, views.Hello())
	})
}
