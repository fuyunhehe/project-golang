package ftime

import (
	"testing"
	"time"
)

func TestFormat1(t *testing.T) {
	cases := []struct {
		In   time.Time
		Want string
	}{
		{In: time.Now(), Want: "2018-11-21"},
	}
	for _, c := range cases {
		got := Format1(c.In)
		if got != c.Want {
			t.Error(c, got)
		}
	}
}

func TestFormat2(t *testing.T) {
	cases := []struct {
		In   time.Time
		Want string
	}{
		{In: time.Now(), Want: "2018-11"},
	}
	for _, c := range cases {
		got := Format2(c.In)
		if got != c.Want {
			t.Error(c, got)
		}
	}
}

func TestFormat3(t *testing.T) {
	cases := []struct {
		In   time.Time
		D    time.Duration
		Want string
	}{
		{In: time.Now(), D: time.Second * 3600, Want: "2018-11-22 11:00:00"},
		{In: time.Now(), D: time.Second * 86400, Want: "2018-11-22 00:00:00"},
	}
	for _, c := range cases {
		got := Format3(c.In, c.D)
		if got != c.Want {
			t.Error(c, got)
		}
		t.Log(c, got)
	}
}

func TestFormat4(t *testing.T)  {
	got := Format4()
	t.Log(got)
}

func TestSub(t *testing.T) {
	cases := []struct {
		Start, End time.Time
		Want       int
	}{
		{Start: time.Now(), End: time.Now().Add(time.Second * 7200), Want: -1},
		{Start: time.Now().Add(time.Second * 7200), End: time.Now(), Want: 1},
		{Start: time.Now(), End: time.Now(), Want: 0},
	}
	for _, c := range cases {
		got := Sub(c.Start, c.End)
		if got == 0 && got != c.Want {
			t.Error(c, got)
		} else if got*c.Want < 0 {
			t.Error(c, got)
		}
	}
}

func TestStr2Time(t *testing.T) {
	cases := []struct {
		S string
	}{
		{S: "0000-00-00 00:00:00"},
		{S: "2018-09-30 12:13:14"},
	}
	for _, c := range cases {
		got, err := Str2Time(c.S)
		t.Log(c.S, got, err)
	}
}
