package main

import (
	"github.com/ezzuodp/go-lib/pkg/fsutils"
	"github.com/ezzuodp/go-lib/pkg/log"
	"github.com/ezzuodp/go-lib/pkg/strutils"
	"github.com/ezzuodp/go-lib/pkg/timeutils"
)

func testFunc() {
	log.Debugf("hello %s", "ionfo")
	log.Warnf("hello %s", "ionfo")
	log.Infof("hello %s", "ionfo")
	log.Errorf("hello %s", "ionfo")

	e := fsutils.WriteFile("/tmp/a.log", strutils.String2Bytes("aaaaabbbbbb\n"), true)
	if e != nil {
		log.Errorf("write file error: %v", e)
	}
}

func main() {
	log.InitConsoleLogger(log.DebugLevel)

	w := timeutils.NewStopwatch("test")
	w.TrackStage("func1", testFunc)
	w.PrintStages()
	w.Reset()
	w.TrackStage("func2", testFunc)
	w.PrintStages()
	w.Reset()
	w.TrackStage("func3", testFunc)
	w.PrintStages()
	w.Reset()
}
