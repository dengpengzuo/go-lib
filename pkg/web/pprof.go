package web

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"net/http/pprof"
)

func innerHandler(h http.HandlerFunc) iris.Handler {
	return func(c iris.Context) {
		h.ServeHTTP(c.ResponseWriter(), c.Request())
	}
}

func IrisDebugHandler(path iris.Party) {
	path.Get("/index", func(c iris.Context) {
		// iris 对 /{prefix}/ => 301 /{prefix}
		// {prefix}/index => {prefix}/
		path := c.Request().URL.Path
		c.Request().URL.Path = path[:len(path)-len("index")]
		pprof.Index(c.ResponseWriter(), c.Request())
	})
	path.Get("/cmdline", innerHandler(pprof.Cmdline))
	path.Get("/profile", innerHandler(pprof.Profile))
	path.Get("/symbol", innerHandler(pprof.Symbol))
	path.Get("/symbol", innerHandler(pprof.Symbol))
	path.Get("/trace", innerHandler(pprof.Trace))
}
