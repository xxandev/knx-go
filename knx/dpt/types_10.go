// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_10001 represents DPT 10.001 / TimeOfDay p. 34.
// Weekday is NOT a golang Weekday, but a KNX Day [0,...,7].
// It may be 0, indicating "no day" was provided.
type DPT_10001 struct {
	Weekday uint8
	Hour    uint8
	Minutes uint8
	Seconds uint8
}

func (d DPT_10001) Pack() []byte {
	var buf = []byte{0, 0, 0, 0}
	if err := d.Valid(); err == nil {
		buf[1] = d.Weekday<<5 | d.Hour&0x1F
		buf[2] = d.Minutes
		buf[3] = d.Seconds
	}
	return []byte(buf)
}

func (d *DPT_10001) Unpack(data []byte) error {
	if len(data) != 4 {
		return ErrInvalidLength
	}
	d.Weekday = uint8(data[1] >> 5)
	d.Hour = uint8(data[1] & 0x1F)
	d.Minutes = uint8(data[2] & 0x3F)
	d.Seconds = uint8(data[3] & 0x3F)
	return d.Valid()
}

func (d DPT_10001) Unit() string {
	return ""
}

func (d DPT_10001) Valid() error {
	if !(d.Weekday <= 7 && d.Hour <= 23 && d.Minutes <= 59 && d.Seconds <= 59) {
		return fmt.Errorf("payload is out of range")
	}
	return nil
}

func (d DPT_10001) String() string {
	weekday := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	if 0 < d.Weekday && d.Weekday <= 7 {
		return fmt.Sprintf("%s %02d:%02d:%02d", weekday[d.Weekday-1], d.Hour, d.Minutes, d.Seconds)
	} else {
		return fmt.Sprintf("%02d:%02d:%02d", d.Hour, d.Minutes, d.Seconds)
	}
}

func (d *DPT_10001) Set(v interface{}) error {
	if data, ok := v.([]uint8); ok {
		if len(data) != 4 {
			return ErrInvalidLength
		}
		d.Weekday = data[0]
		d.Hour = data[1]
		d.Minutes = data[2]
		d.Seconds = data[3]
		return d.Valid()
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_10001) Get() interface{} {
	return []uint8{d.Weekday, d.Hour, d.Minutes, d.Seconds}
}

func (d DPT_10001) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, d.Weekday, d.Hour, d.Minutes, d.Seconds)
	}
	return fmt.Sprintf("%d %02d:%02d:%02d", d.Weekday, d.Hour, d.Minutes, d.Seconds)
}
