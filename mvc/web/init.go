package web

import (
	"blog/mvc/zz"
	"blog/mvc/zz/ctx"
	"blog/system"
	"blog/tools"
	"blog/tools/conf"
	"html/template"
	"io/ioutil"
	"path"
)

func Str2html(raw string) template.HTML {
	return template.HTML(raw)
}

var SysConf = func(c *ctx.Context) *zz.InterceptErr {
	c.AddHtmlParam("navs", conf.Nav.Navs())
	c.AddHtmlParam("base", system.Conf.Server.Base)
	return zz.NewInterceptOK()
}

func init() {
	zz.AddInterceptor(SysConf)
	zz.Engine.SetFuncMap(map[string]interface{}{
		"str2html": Str2html,
	})
	zz.Engine.Static(tools.BaseAdd("/static"), system.Conf.Server.Static)
	zz.Engine.LoadHTMLGlob(system.Conf.Server.HTML)

	zz.GetMapping(tools.BaseAdd("/index"), func(c *ctx.Context) interface{} {
		nav := c.Query("nav")
		var links []*conf.BlogInfo
		if nav == "" {
			links = conf.AllBlog[:5]
		} else {
			links = conf.Nav.Links(nav)
		}
		c.AddHtmlParam("links", links)
		return "Aindex"
	})
	zz.GetMapping(tools.BaseAdd("/archive"), func(c *ctx.Context) interface{} {
		c.AddHtmlParam("yLinks", conf.YearBlog)
		return "Aarchive"
	})
	zz.GetMapping(tools.BaseAdd("/content/:nav/:file"), func(c *ctx.Context) interface{} {
		nav := c.Param("nav")
		filename := c.Param("file")
		file, err := ioutil.ReadFile(path.Join(conf.Nav.Dir, nav, filename))
		if err != nil {
			panic(err)
		}
		c.AddHtmlParam("navStr", nav)
		c.AddHtmlParam("content", string(file))
		c.AddHtmlParam("title", tools.PathToName(filename))
		return "Acontent"
	})

	zz.GetMapping(tools.BaseAdd("/img/:nav/_image/:path/:file"), func(c *ctx.Context) interface{} {
		pathS := c.Param("path")
		nav := c.Param("nav")
		file := c.Param("file")
		return zz.HTTP_IMG + path.Join(conf.Nav.Dir, nav, "_image", pathS, file)
	})
}
