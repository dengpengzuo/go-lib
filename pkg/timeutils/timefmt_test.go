package timeutils

import (
	"testing"
	"time"
)

func TestFmtDate(t *testing.T) {
	now := time.Now()
	t.Log(FmtDate(now))
	t.Log(FmtDateTime(now))
	t.Log(FmtISO8601DateTime(now))
	t.Log(FmtLogDateTime(now))
}
