package ctx

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func (c *Context) SetHtmlTemp(data interface{}) {
	c.Set("temp", data)
}

func (c *Context) AddHtmlParam(key string, data interface{}) {
	if _, exists := c.Get("html"); !exists {
		c.Set("html", make(map[string]interface{}))
	}
	get, _ := c.Get("html")
	m := get.(map[string]interface{})
	m[key] = data
}
