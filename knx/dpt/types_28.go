// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import "fmt"

// DPT_28001 represents DPT 28.001 / Var String UTF-8.
type DPT_28001 string

func (d DPT_28001) Pack() []byte {
	// len(d) is gives us the number of bytes in d
	var buf = make([]byte, 1, len(d)+2)

	buf = append(buf, d...)
	buf = append(buf, 0x00)

	return buf
}

func (d *DPT_28001) Unpack(data []byte) error {
	var buf = data[1 : len(data)-1]

	*d = DPT_28001(buf)

	return nil
}

func (d DPT_28001) Unit() string {
	return ""
}

func (d DPT_28001) String() string {
	return string(d)
}

func (d *DPT_28001) Set(v interface{}) error {
	if data, ok := v.(string); ok {
		*d = DPT_28001(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_28001) Get() interface{} {
	return string(d)
}

func (d DPT_28001) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, string(d))
	}
	return fmt.Sprint(string(d))
}
