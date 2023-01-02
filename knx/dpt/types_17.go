package dpt

import "fmt"

// DPT_17001 represents DPT 17.001 / Scene Number.
type DPT_17001 uint8

func (d DPT_17001) Pack() []byte {
	if d > 63 {
		return packU8(63)
	} else {
		return packU8(uint8(d))
	}
}

func (d *DPT_17001) Unpack(data []byte) error {
	var value uint8

	if err := unpackU8(data, &value); err != nil {
		return err
	}

	if value <= 63 {
		*d = DPT_17001(value)
		return nil
	} else {
		*d = DPT_17001(63)
		return nil
	}
}

func (d DPT_17001) Unit() string {
	return ""
}

func (d DPT_17001) String() string {
	return fmt.Sprintf("%d", uint8(d))
}

func (d *DPT_17001) Set(v interface{}) error {
	switch data := v.(type) {
	case int8:
		*d = DPT_17001(data)
		return nil
	case int:
		*d = DPT_17001(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_17001) Get() interface{} {
	return int8(d)
}

func (d DPT_17001) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, int8(d))
	}
	return fmt.Sprint(int8(d))
}
