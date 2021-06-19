package timeutils

import (
	"github.com/ezzuodp/go-lib/pkg/log"
	"github.com/ezzuodp/go-lib/pkg/strutils"
	"testing"
	"time"
)

func TestFmtDate(t *testing.T) {
	s := "aaaaaaa"
	bytes := []byte(s)
	v := strutils.Bytes2String(bytes[:])
	log.Infof("ssss=>[%s]", v)

	now := time.Now()
	log.Infof("%d", NowUnix())
	log.Infof("%d", NowMillis())
	log.Infof("%s", FmtDate(now))
	log.Infof("%s", FmtDateTime(now))
	log.Infof("%s", FmtISO8601DateTime(now))
	log.Infof("%s", FmtLogDateTime(now))
}
