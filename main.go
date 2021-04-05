package main

import (
	"context"
	"github.com/coreos/etcd/embed"
	"github.com/ezzuodp/go-lib/pkg/emetcd"
	"github.com/ezzuodp/go-lib/pkg/log"
	"github.com/ezzuodp/go-lib/pkg/proc"
	"github.com/ezzuodp/go-lib/pkg/web"
	"github.com/kataras/iris/v12"
	context2 "github.com/kataras/iris/v12/context"
	"net/http"
	"os"
	"time"
)

const (
	address = ":8080"
)

type Server struct {
	http *http.Server

	app  *iris.Application
	etcd *embed.Etcd
}

var s *Server

func closeServer() {
	if s != nil {
		s.app.Shutdown(context.Background())
		s.etcd.Close()
	}
}

func init() {
	log.InitConsoleLogger(log.DebugLevel)
	proc.AddShutdownHook(closeServer)
}

func main() {
	proc.GoWatchSignal()

	s = &Server{}
	s.app = web.NewIris(
		web.IrisLogger(time.RFC3339, "info"),
		web.IrisAccessLogger(),
		// request router path config
		web.IrisAddPathHandler("/debug/pprof", web.IrisDebugHandler),

		web.IrisGetHandler("/ping", func(ctx context2.Context) {
			ctx.JSON(&struct {
				Code int32
				Msg  string
			}{Code: 0, Msg: "成功"})
		}),
	)

	s.http = &http.Server{
		Addr: address,
	}

	s.etcd, _ = emetcd.GenEmbedEtcd()
	err := web.StartServe(s.app, s.http)
	if err != nil {
		log.Errorf("web serve error: %v", err)
		os.Exit(-1)
	}
}
