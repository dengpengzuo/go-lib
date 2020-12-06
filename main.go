package main

import (
    "github.com/go-lib/pkg/fsutils"
    "github.com/go-lib/pkg/log"
    "github.com/go-lib/pkg/strutils"
)

func main() {
    log.InitConsoleLogger(log.DebugLevel)

    log.Debugf("infof ====> %s", "ionfo")
    log.Warnf("infof ====> %s", "ionfo")
    log.Infof("infof ====> %s", "ionfo")
    log.Errorf("infof ====> %s", "ionfo")

    e := fsutils.WriteFile("/tmp/a.log", strutils.String2Bytes("aaaaabbbbbb\n"), true)
    if e != nil {
        log.Errorf("write file error: %v", e)
    }
}
