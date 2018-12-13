package ftime

import (
	"time"
)

func Format(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func Format1(t time.Time) string {
	return t.Format("2006-01-02")
}

func Format2(t time.Time) string {
	return t.Format("2006-01")
}

func Format3(t time.Time, d time.Duration) string {
	return t.Truncate(d).Format("2006-01-02 15:04:05")
}

func Format4() string {
	return time.Date(2000, 3, 0, 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
}

func Sub(start, end time.Time) int {
	return int(start.Sub(end).Hours())
}

func Str2Time(s string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", s)
}
