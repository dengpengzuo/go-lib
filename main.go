package main

import (
	"github.com/ezzuodp/go-lib/pkg/log"
	"github.com/ezzuodp/go-lib/pkg/web"
	"net/http"
	"os"
	"time"
)

const (
	address = ":8080"
)

func main() {
	log.InitConsoleLogger(log.DebugLevel)

	app := web.InitIris(
		web.IrisLogger(time.RFC3339, "debug"),
		web.IrisAccessLogger(),
		// request router path config
		web.IrisAddPathHandler("/debug/pprof", web.IrisDebugHandler),
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
