package tinytime_test

import (
	"testing"
	"time"

	"github.com/go-compile/tinytime"
)

func TestCreateFromUnix(t *testing.T) {
	now := time.Now()
	s := tinytime.Unix(now.Unix())

	if s.Seconds() != now.Unix() {
		t.Error("returned incorrect unix time")
	}
}

func TestToBytes(t *testing.T) {
	now := time.Now()
	s := tinytime.Unix(now.Unix())

	if tinytime.FromBytes(
		tinytime.ToBytes(s.Minutes()),
	) != s.Minutes() {
		t.Error("returned incorrect time")
	}
}

func TestToFixed16bit(t *testing.T) {
	now := time.Now()
	s := tinytime.Unix(now.Unix())

	buf := tinytime.ToFixed(tinytime.ToBytes(s.Days()), 2)

	if len(buf) != 2 {
		t.Error("buffer does not equal fixed length")
	}

	if tinytime.FromBytes(
		buf,
	) != s.Days() {
		t.Error("returned incorrect time")
	}
}

func TestToFixed32bit(t *testing.T) {
	now := time.Now()
	s := tinytime.Unix(now.Unix())

	buf := tinytime.ToFixed(tinytime.ToBytes(s.Minutes()), 4)

	if len(buf) != 4 {
		t.Error("buffer does not equal fixed length")
	}

	if tinytime.FromBytes(
		buf,
	) != s.Minutes() {
		t.Error("returned incorrect time")
	}
}

func TestToFixed64bit(t *testing.T) {
	now := time.Now()
	s := tinytime.Unix(now.Unix())

	buf := tinytime.ToFixed(tinytime.ToBytes(s.Minutes()), 8)

	if len(buf) != 8 {
		t.Error("buffer does not equal fixed length")
	}

	if tinytime.FromBytes(
		buf,
	) != s.Minutes() {
		t.Error("returned incorrect time")
	}
}

func TestMinutes(t *testing.T) {
	timeStamp := int64(30720) //512 minutes
	s := tinytime.Unix(timeStamp)

	if s.Minutes() != 512 {
		t.Error("returned incorrect minutes")
	}
}

func TestHours(t *testing.T) {
	timeStamp := int64(10800) //3 hours
	s := tinytime.Unix(timeStamp)

	if s.Hours() != 3 {
		t.Error("returned incorrect hours")
	}
}

func TestTwoHours(t *testing.T) {
	timeStamp := int64(86400) //24 hours
	s := tinytime.Unix(timeStamp)

	if s.TwoHours() != 12 {
		t.Error("returned incorrect twoHours")
	}
}

func TestThreeHours(t *testing.T) {
	timeStamp := int64(86400) //24 hours
	s := tinytime.Unix(timeStamp)

	if s.ThreeHours() != 8 {
		t.Error("returned incorrect threeHours")
	}
}

func TestFourHours(t *testing.T) {
	timeStamp := int64(86400) //24 hours
	s := tinytime.Unix(timeStamp)

	if s.FourHours() != 6 {
		t.Error("returned incorrect fourHours")
	}
}

func TestFiveHours(t *testing.T) {
	timeStamp := int64(86400) //24 hours
	s := tinytime.Unix(timeStamp)

	if s.FiveHours() != 4 {
		t.Error("returned incorrect fiveHours")
	}
}

func TestDays(t *testing.T) {
	timeStamp := int64(5529600) //64 days
	s := tinytime.Unix(timeStamp)

	if s.Days() != 64 {
		t.Error("returned incorrect days")
	}
}
