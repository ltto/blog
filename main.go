package main

import (
	_ "blog/mvc/web"
	"blog/system"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ltto/T/www"
)

func main() {
	gin.SetMode("release")
	if err := www.Engine.Run(fmt.Sprintf("127.0.0.1:%d", system.Conf.Server.Port)); err != nil {
		panic(err)
	}
}
