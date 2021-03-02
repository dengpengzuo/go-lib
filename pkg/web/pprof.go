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
		// iris对 /debug/pprof/ => 301 /debug/pprof
		// 但 pprof.index 方法中判断的URL.PATH必须是 /debug/pprof/
		c.Request().URL.Path = "/debug/pprof/"
		pprof.Index(c.ResponseWriter(), c.Request())
	})
	path.Get("/cmdline", innerHandler(pprof.Cmdline))
	path.Get("/profile", innerHandler(pprof.Profile))
	path.Get("/symbol", innerHandler(pprof.Symbol))
	path.Get("/symbol", innerHandler(pprof.Symbol))
	path.Get("/trace", innerHandler(pprof.Trace))
}
