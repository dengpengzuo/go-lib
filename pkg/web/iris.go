package web

import (
	"github.com/ezzuodp/go-lib/pkg/log"
	"github.com/ezzuodp/go-lib/pkg/tcp"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	ircover "github.com/kataras/iris/v12/middleware/recover"
	"net"
	"net/http"
	"time"
)

type IrisOptionFunc func(*iris.Application)

func InitIris(options ...IrisOptionFunc) *iris.Application {
	app := iris.New()
	for _, option := range options {
		option(app)
	}
	return app
}

func IrisAccessLogger() IrisOptionFunc {
	return func(app *iris.Application) {
		accessLogger := logger.New(logger.Config{
			Status: true,
			IP:     true,
			Method: true,
			Path:   true,
			Query:  true,

			// 若为空则从 `ctx.Values().Get("logger_message") 获取内容
			// 在其中添加日志
			MessageContextKeys: []string{"logger_message"},

			// 为空则从 `ctx.GetHeader("User-Agent") 获取头信息
			MessageHeaderKeys: []string{"User-Agent"},

			LogFunc: func(endTime time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
				// 127.0.0.1 - - [21/Sep/2020:11:46:07 +0800] "GET /api/hello HTTP/1.1" 404 283 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36"
				log.Infof("%v - - [%s] \"%s %s\" %s %v \"%v\"", ip, endTime.Format(time.RFC3339), method, path, status, latency, headerMessage)
			},
		})
		app.Use(ircover.New(), accessLogger)
	}
}

func IrisLogger(timefmt, level string) IrisOptionFunc {
	return func(app *iris.Application) {
		app.Logger().SetTimeFormat(timefmt) // time.RFC3339
		app.Logger().SetLevel(level)        // iris 自身logger
	}
}

type IrisPathPartyFunc func(iris.Party)

func IrisAddPathHandler(path string, options ...IrisPathPartyFunc) IrisOptionFunc {
	return func(app *iris.Application) {
		party := app.Party(path)
		for _, o := range options {
			o(party)
		}
	}
}

func StartServe(app *iris.Application, server *http.Server) error {
	var err error

	err = app.Build()
	if err != nil {
		return err
	}

	var listener net.Listener
	if listener, err = tcp.Listen("tcp", server.Addr); err != nil {
		log.Errorf("listen address error: %v", err)
		return err
	}

	log.Infof("web server listen %s ... \n", server.Addr)
	return app.NewHost(server).Serve(listener)
}
