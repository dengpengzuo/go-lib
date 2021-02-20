package timeutils

import (
	"github.com/ezzuodp/go-lib/pkg/strutils"
	"testing"
	"time"
)

func TestFmtDate(t *testing.T) {
	s := "aaaaaaa"
	bytes := []byte(s)
	v := strutils.Bytes2String(bytes[:])
	t.Logf("ssss=>[%s]", v)
	now := time.Now()
	t.Log(NowUnix())
	t.Log(NowMillis())
	t.Log(FmtDate(now))
	t.Log(FmtDateTime(now))
	t.Log(FmtISO8601DateTime(now))
	t.Log(FmtLogDateTime(now))
}
