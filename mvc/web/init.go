package web

import (
	"blog/system"
	"blog/tools"
	"blog/tools/conf"
	"html/template"
	"io/ioutil"
	"path"

	"github.com/ltto/T/www"
)

func Str2html(raw string) template.HTML {
	return template.HTML(raw)
}

var SysConf = func(c *www.Context) *www.InterceptErr {
	c.AddCParam("navs", conf.Nav.Navs())
	c.AddCParam("base", system.Conf.Server.Base)
	return www.NewInterceptOK()
}

func init() {
	www.AddInterceptor(SysConf)
	www.Engine.SetFuncMap(map[string]interface{}{
		"str2html": Str2html,
	})
	www.Engine.Static(tools.BaseAdd("/static"), system.Conf.Server.Static)
	www.Engine.LoadHTMLGlob(system.Conf.Server.HTML)

	www.GetMapping(tools.BaseAdd("/index"), func(c *www.Context) interface{} {
		nav := c.Query("nav")
		var links []*conf.BlogInfo
		if nav == "" {
			links = conf.AllBlog[:5]
		} else {
			links = conf.Nav.Links(nav)
		}
		c.AddCParam("links", links)
		return "Aindex.html"
	})
	www.GetMapping(tools.BaseAdd("/archive"), func(c *www.Context) interface{} {
		c.AddCParam("yLinks", conf.YearBlog)
		return "Aarchive.html"
	})
	www.GetMapping(tools.BaseAdd("/content/:nav/:file"), func(c *www.Context) interface{} {
		nav := c.Param("nav")
		filename := c.Param("file")
		file, err := ioutil.ReadFile(path.Join(conf.Nav.Dir, nav, filename))
		if err != nil {
			panic(err)
		}
		c.AddCParam("navStr", nav)
		c.AddCParam("content", string(file))
		c.AddCParam("title", tools.PathToName(filename))
		return "Acontent.html"
	})

	www.GetMapping(tools.BaseAdd("/img/:nav/_image/:path/:file"), func(c *www.Context) interface{} {
		pathS := c.Param("path")
		nav := c.Param("nav")
		file := c.Param("file")
		return www.HttpImg + path.Join(conf.Nav.Dir, nav, "_image", pathS, file)
	})
}
