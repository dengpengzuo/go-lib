package proc

import (
	"github.com/ezzuodp/go-lib/pkg/log"
	"os"
	"os/signal"
	"syscall"
)

type HookFunc func()

var gHooks []HookFunc

func init() {
	gHooks = make([]HookFunc, 0, 16)
}

func AddShutdownHook(hook HookFunc) {
	gHooks = append(gHooks, hook)
}

func GoWatchSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-c
		log.Warnf("receive signal %v, call shutdown hook ... \n", s)
		for i := len(gHooks) - 1; i >= 0; i-- {
			gHooks[i]()
		}
		os.Exit(1) // second signal. Exit directly.
	}()
}
