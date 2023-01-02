package dpt

import "fmt"

// DPT_12001 represents DPT 12.001 / Unsigned counter.
type DPT_12001 uint32

func (d DPT_12001) Pack() []byte {
	return packU32(uint32(d))
}

func (d *DPT_12001) Unpack(data []byte) error {
	return unpackU32(data, (*uint32)(d))
}

func (d DPT_12001) Unit() string {
	return "pulses"
}

func (d DPT_12001) String() string {
	return fmt.Sprintf("%d pulses", uint32(d))
}

func (d *DPT_12001) Set(v interface{}) error {
	switch data := v.(type) {
	case uint32:
		*d = DPT_12001(data)
		return nil
	case int:
		*d = DPT_12001(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_12001) Get() interface{} {
	return uint32(d)
}

func (d DPT_12001) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, uint32(d))
	}
	return fmt.Sprintf("%d", uint32(d))
}
