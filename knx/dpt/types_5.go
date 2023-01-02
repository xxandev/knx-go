package dpt

import (
	"fmt"
)

// DPT_5001 represents DPT 5.001 / Scaling.
type DPT_5001 float32

func (d DPT_5001) Pack() []byte {
	if d <= 0 {
		return packU8(0)
	} else if d >= 100 {
		return packU8(255)
	} else {
		return packU8(uint8(d * 2.55))
	}
}

func (d *DPT_5001) Unpack(data []byte) error {
	var value uint8
	if err := unpackU8(data, &value); err != nil {
		return err
	}

	*d = DPT_5001(value) / 2.55

	return nil
}

func (d DPT_5001) Unit() string {
	return "%"
}

func (d DPT_5001) String() string {
	return fmt.Sprintf("%.2f%%", float32(d))
}

func (d *DPT_5001) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_5001(data)
		return nil
	case float64:
		*d = DPT_5001(data)
		return nil
	case int:
		*d = DPT_5001(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_5001) Get() interface{} {
	return float32(d)
}

func (d DPT_5001) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_5003 represents DPT 5.003 / Angle.
type DPT_5003 float32

func (d DPT_5003) Pack() []byte {
	if d <= 0 {
		return packU8(0)
	} else if d >= 360 {
		return packU8(255)
	} else {
		return packU8(uint8(d * 255 / 360))
	}
}

func (d *DPT_5003) Unpack(data []byte) error {
	var value uint8
	if err := unpackU8(data, &value); err != nil {
		return err
	}

	*d = DPT_5003(value) * 360 / 255

	return nil
}

func (d DPT_5003) Unit() string {
	return "°"
}

func (d DPT_5003) String() string {
	return fmt.Sprintf("%.2f°", float32(d))
}

func (d *DPT_5003) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_5003(data)
		return nil
	case float64:
		*d = DPT_5003(data)
		return nil
	case int:
		*d = DPT_5003(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_5003) Get() interface{} {
	return float32(d)
}

func (d DPT_5003) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_5004 represents DPT 5.004 / Percent_U8.
type DPT_5004 uint8

func (d DPT_5004) Pack() []byte {
	return packU8(uint8(d))
}

func (d *DPT_5004) Unpack(data []byte) error {
	return unpackU8(data, (*uint8)(d))
}

func (d DPT_5004) Unit() string {
	return "%"
}

func (d DPT_5004) String() string {
	return fmt.Sprintf("%.2f%%", float32(d))
}

func (d *DPT_5004) Set(v interface{}) error {
	switch data := v.(type) {
	case uint8:
		*d = DPT_5004(data)
		return nil
	case int:
		*d = DPT_5004(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_5004) Get() interface{} {
	return uint8(d)
}

func (d DPT_5004) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, uint8(d))
	}
	return fmt.Sprintf("%d", uint8(d))
}

// DPT_5005 represents DPT 5.005 / Ratio (0..255).
type DPT_5005 uint8

func (d DPT_5005) Pack() []byte {
	return packU8(uint8(d))
}

func (d *DPT_5005) Unpack(data []byte) error {
	return unpackU8(data, (*uint8)(d))
}

func (d DPT_5005) Unit() string {
	return ""
}

func (d DPT_5005) String() string {
	return fmt.Sprintf("%d", uint8(d))
}

func (d *DPT_5005) Set(v interface{}) error {
	switch data := v.(type) {
	case uint8:
		*d = DPT_5005(data)
		return nil
	case int:
		*d = DPT_5005(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_5005) Get() interface{} {
	return uint8(d)
}

func (d DPT_5005) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, uint8(d))
	}
	return fmt.Sprintf("%d", uint8(d))
}
