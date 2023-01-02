// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

// A DatapointValue is a value of a datapoint.
type DatapointValue interface {
	// Pack the datapoint to a byte array.
	Pack() []byte

	// Unpack a the datapoint value from a byte array.
	Unpack(data []byte) error

	// Set - set DPT_xxx value
	Set(v interface{}) error

	// Get - get DPT_xxx value
	Get() interface{}

	// Formatting - return value to a string of the given format
	//  Formatting("") return default format value
	//  Formatting("%v") return format value
	//  ...
	Formatting(format string) string
}

// DatapointMeta gives meta information about a datapoint type.
type DatapointMeta interface {
	// Unit returns the unit of this datapoint type or empty string if it doesn't have a unit.
	Unit() string
}
