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

func ping(ctx iris.Context) {
	ctx.JSON(&JsonRsp{Code: 0, Msg: "成功"})
}

func main() {
	log.InitConsoleLogger(log.DebugLevel)

	app := web.NewIris(
		web.IrisLogger(time.RFC3339, "info"),
		web.IrisAccessLogger(),
		// request router path config
		web.IrisAddPathHandler("/debug/pprof", web.IrisDebugHandler),
		web.IrisGetHandler("/ping", ping),
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
