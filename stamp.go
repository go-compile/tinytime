package unixtime

import (
	"time"
)

type Stamp int64

func Now() Stamp {
	return Stamp(time.Now().Unix())
}

func Unix(time int64) Stamp {
	return Stamp(time)
}

func (s Stamp) Seconds() int64 {
	return int64(s)
}

func (s Stamp) Minutes() int64 {
	return int64(s / 60)
}

func (s Stamp) Hours() int64 {
	return int64((s / 60) / 60)
}

func (s Stamp) TwoHours() int64 {
	return int64((s/60)/60) / 2
}

func (s Stamp) ThreeHours() int64 {
	return int64((s/60)/60) / 3
}

func (s Stamp) FourHours() int64 {
	return int64((s/60)/60) / 4
}

func (s Stamp) FiveHours() int64 {
	return int64((s/60)/60) / 5
}
