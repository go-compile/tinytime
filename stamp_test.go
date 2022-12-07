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

func TestTwoHours(t *testing.T) {
	timeStamp := int64(86400) //24 hours
	s := unixtime.Unix(timeStamp)

	if s.TwoHours() != 12 {
		t.Error("returned incorrect twoHours")
	}
}

func TestThreeHours(t *testing.T) {
	timeStamp := int64(86400) //24 hours
	s := unixtime.Unix(timeStamp)

	if s.ThreeHours() != 8 {
		t.Error("returned incorrect threeHours")
	}
}

func TestFourHours(t *testing.T) {
	timeStamp := int64(86400) //24 hours
	s := unixtime.Unix(timeStamp)

	if s.FourHours() != 6 {
		t.Error("returned incorrect fourHours")
	}
}

func TestFiveHours(t *testing.T) {
	timeStamp := int64(86400) //24 hours
	s := unixtime.Unix(timeStamp)

	if s.FiveHours() != 4 {
		t.Error("returned incorrect fiveHours")
	}
}

func TestDays(t *testing.T) {
	timeStamp := int64(5529600) //64 days
	s := unixtime.Unix(timeStamp)

	if s.Days() != 64 {
		t.Error("returned incorrect days")
	}
}
