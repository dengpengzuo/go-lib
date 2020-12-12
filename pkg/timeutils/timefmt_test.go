package timeutils

import (
	"testing"
	"time"
)

func TestFmtDate(t *testing.T) {
	now := time.Now()
	t.Log(NowUnix())
	t.Log(NowMillis())
	t.Log(FmtDate(now))
	t.Log(FmtDateTime(now))
	t.Log(FmtISO8601DateTime(now))
	t.Log(FmtLogDateTime(now))
}
