package unixtime_test

import (
	"testing"
	"time"

	"github.com/go-compile/unixtime"
)

func TestCreateFromUnix(t *testing.T) {
	now := time.Now()
	s := unixtime.Unix(now.Unix())

	if s.Seconds() != now.Unix() {
		t.Error("returned incorrect unix time")
	}
}

func TestMinutes(t *testing.T) {
	timeStamp := int64(30720) //512 minutes
	s := unixtime.Unix(timeStamp)

	if s.Minutes() != 512 {
		t.Error("returned incorrect minutes")
	}
}

func TestHours(t *testing.T) {
	timeStamp := int64(10800) //3 hours
	s := unixtime.Unix(timeStamp)

	if s.Hours() != 3 {
		t.Error("returned incorrect hours")
	}
}
