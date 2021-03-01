package timeutils

import "time"

func NowUnix() int64 {
	return time.Now().Unix()
}

func NowMillis() int64 {
	v := time.Now().UnixNano()
	return v / int64(time.Millisecond)
}

func MillisecondToTime(v int64) time.Time {
	return time.Unix(0, v*int64(time.Millisecond))
}

func NanosecondToTime(v int64) time.Time {
	return time.Unix(0, v)
}

func FmtDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func FmtNowDate() string {
	return time.Now().Format("2006-01-02")
}

func ParseDate(v string) (time.Time, error) {
	return time.Parse("2006-01-02", v)
}

func FmtDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func FmtNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ParseDateTime(v string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", v)
}

func FmtISO8601DateTime(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000Z0700")
}

func FmtNowISO8601DateTime() string {
	return time.Now().Format("2006-01-02T15:04:05.000Z0700")
}

func FmtLogDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.000000")
}

func FmtLogNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05.000000")
}
