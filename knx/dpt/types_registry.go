// Copyright 2020 Sven Rebhan.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"reflect"
	"sync"
)

var (
	once     sync.Once
	registry map[string]reflect.Type
)

// Init function used to add all types
func setup() {
	// Singleton, can only run once
	once.Do(func() {
		// Register the types
		registry = make(map[string]reflect.Type)
		for _, d := range GetTypes() {
			// Determine the name of the datatype
			d_type := reflect.TypeOf(d).Elem()
			name := d_type.Name()

			// Convert the name into KNX yy.xxx (e.g. DPT_1001 --> 1.001) format
			name = name[4:len(name)-3] + "." + name[len(name)-3:]

			// Register the type
			registry[name] = d_type
		}
	})
}

// ListSupportedTypes returns the name all known datapoint-types (DPTs).
func ListSupportedTypes() []string {
	// Setup the registry
	setup()

	// Initialize the key-list
	keys := make([]string, len(registry))

	// Fill the key-list
	i := 0
	for k := range registry {
		keys[i] = k
		i++
	}

	return keys
}

// Produce creates a new instance of the given datapoint-type name e.g. "1.001".
func Produce(name string) (d DatapointValue, ok bool) {
	// Setup the registry
	setup()

	// Lookup the given type and create a new instance of that type
	x, ok := registry[name]
	if ok {
		d = reflect.New(x).Interface().(DatapointValue)
	}
	return d, ok
}
