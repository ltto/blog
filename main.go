package main

import (
	_ "blog/mvc/web"
	"blog/mvc/zz"
	"blog/system"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release")
	if err := zz.Engine.Run(fmt.Sprintf("127.0.0.1:%d", system.Conf.Server.Port)); err != nil {
		panic(err)
	}
}
