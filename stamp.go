package tinytime

import (
	"encoding/binary"
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
	return int64(s / 3600)
}

func (s Stamp) TwoHours() int64 {
	return int64(s/3600) / 2
}

func (s Stamp) ThreeHours() int64 {
	return int64(s/3600) / 3
}

func (s Stamp) FourHours() int64 {
	return int64(s/3600) / 4
}

func (s Stamp) FiveHours() int64 {
	return int64(s/3600) / 5
}

func (s Stamp) Days() int64 {
	return int64(s/3600) / 24
}

// ToBytes will take a timestamp and format it as small as possible
func ToBytes(timestamp int64) []byte {
	buf := make([]byte, 8)

	// right align bytes
	binary.BigEndian.PutUint64(buf, uint64(timestamp))

	// loop for all bytes other than the last
	for i := 0; i < 7; i++ {

		// if byte empty delete it
		if buf[0] == 0 {
			buf = buf[1:]
			continue
		}

		// if byte not empty break because we have hit the timestamp
		break
	}

	return buf
}

// FromBytes decodes a int of any size uint8 to int64
func FromBytes(buf []byte) int64 {
	// if buf less than 8 bytes add the appropriate amount to
	// bring it back to 8
	pad := make([]byte, 8-len(buf))
	buf = append(pad, buf...)

	return int64(binary.BigEndian.Uint64(buf))
}

// ToFixed takes the output from func ToBytes() then either extends or truncates
// it to fixed the desired length.
// If you provide a buffer larger than length this function will panic to prevent
// data loss.
func ToFixed(buf []byte, length int) []byte {
	if len(buf) > length {
		panic("you cannot delete bytes which contain data")
	}

	// append needed bytes to the start
	pad := make([]byte, length-len(buf))
	return append(pad, buf...)
}
