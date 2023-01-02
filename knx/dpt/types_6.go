package dpt

import (
	"fmt"
)

// DPT_6010 represents DPT 6.010 / counter pulses (-128..127).
type DPT_6010 int8

func (d DPT_6010) Pack() []byte {
	return packV8(int8(d))
}

func (d *DPT_6010) Unpack(data []byte) error {
	return unpackV8(data, (*int8)(d))
}

func (d DPT_6010) Unit() string {
	return "counter pulses"
}

func (d DPT_6010) String() string {
	return fmt.Sprintf("%d counter pulses", int8(d))
}

func (d *DPT_6010) Set(v interface{}) error {
	switch data := v.(type) {
	case int8:
		*d = DPT_6010(data)
		return nil
	case int:
		*d = DPT_6010(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_6010) Get() interface{} {
	return int8(d)
}

func (d DPT_6010) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, int8(d))
	}
	return fmt.Sprint(int8(d))
}
