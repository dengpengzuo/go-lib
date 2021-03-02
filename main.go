package main

import (
	"github.com/ezzuodp/go-lib/pkg/log"
	"github.com/ezzuodp/go-lib/pkg/web"
	"github.com/kataras/iris/v12"
	"net/http"
	"os"
	"time"
)

const (
	address = ":8080"
)

type JsonRsp struct {
	Code int32
	Msg  string
}

func hello(ctx iris.Context) {
	ctx.JSON(&JsonRsp{Code: 0, Msg: "成功"})
}

func main() {
	log.InitConsoleLogger(log.DebugLevel)

	app := web.InitIris(
		web.IrisLogger(time.RFC3339, "debug"),
		web.IrisAccessLogger(),
		// request router path config
		web.IrisAddPathHandler("/debug/pprof", web.IrisDebugHandler),
		web.IrisGetHandler("/hello", hello),
	)

	server := &http.Server{
		Addr: address,
	}

	err := web.StartServe(app, server)
	if err != nil {
		log.Errorf("web serve error: %v", err)
		os.Exit(-1)
	}
}
