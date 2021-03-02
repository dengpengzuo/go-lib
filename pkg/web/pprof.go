package web

import (
	"github.com/kataras/iris/v12"
	"net/http/pprof"
)

func IrisDebugHandler(path iris.Party) {
	path.Get("/index", func(c iris.Context) {
		// pprof.Index 只认格式: {prefix}/, 所以将 {prefix}/index 后 index 删除
		path := c.Request().URL.Path
		c.Request().URL.Path = path[:len(path)-len("index")]
		pprof.Index(c.ResponseWriter(), c.Request())
	})
	path.Get("/cmdline", HttpHandlerAdapter(pprof.Cmdline))
	path.Get("/profile", HttpHandlerAdapter(pprof.Profile))
	path.Get("/symbol", HttpHandlerAdapter(pprof.Symbol))
	path.Get("/trace", HttpHandlerAdapter(pprof.Trace))
}
