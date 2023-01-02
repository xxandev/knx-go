// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_14000 represents DPT 14.000 / Acceleration
type DPT_14000 float32

func (d DPT_14000) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14000) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14000(value)

	return nil
}

func (d DPT_14000) Unit() string {
	return "m/s²"
}

func (d DPT_14000) String() string {
	return fmt.Sprintf("%.2f m/s²", float32(d))
}

func (d *DPT_14000) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14000(data)
		return nil
	case float64:
		*d = DPT_14000(data)
		return nil
	case int:
		*d = DPT_14000(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14000) Get() interface{} {
	return float32(d)
}

func (d DPT_14000) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14001 represents DPT 14.001 / Acceleration Angular
type DPT_14001 float32

func (d DPT_14001) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14001) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14001(value)

	return nil
}

func (d DPT_14001) Unit() string {
	return "rad/s²"
}

func (d DPT_14001) String() string {
	return fmt.Sprintf("%.2f rad/s²", float32(d))
}

func (d *DPT_14001) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14001(data)
		return nil
	case float64:
		*d = DPT_14001(data)
		return nil
	case int:
		*d = DPT_14001(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14001) Get() interface{} {
	return float32(d)
}

func (d DPT_14001) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14002 represents DPT 14.002 / ActivationEnergy
type DPT_14002 float32

func (d DPT_14002) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14002) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14002(value)

	return nil
}

func (d DPT_14002) Unit() string {
	return "J/mol"
}

func (d DPT_14002) String() string {
	return fmt.Sprintf("%.2f J/mol", float32(d))
}

func (d *DPT_14002) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14002(data)
		return nil
	case float64:
		*d = DPT_14002(data)
		return nil
	case int:
		*d = DPT_14002(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14002) Get() interface{} {
	return float32(d)
}

func (d DPT_14002) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14003 represents DPT 14.003 / Activity
type DPT_14003 float32

func (d DPT_14003) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14003) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14003(value)

	return nil
}

func (d DPT_14003) Unit() string {
	return "s⁻¹"
}

func (d DPT_14003) String() string {
	return fmt.Sprintf("%.2f s⁻¹", float32(d))
}

func (d *DPT_14003) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14003(data)
		return nil
	case float64:
		*d = DPT_14003(data)
		return nil
	case int:
		*d = DPT_14003(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14003) Get() interface{} {
	return float32(d)
}

func (d DPT_14003) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14004 represents DPT 14.004 / Mol
type DPT_14004 float32

func (d DPT_14004) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14004) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14004(value)

	return nil
}

func (d DPT_14004) Unit() string {
	return "mol"
}

func (d DPT_14004) String() string {
	return fmt.Sprintf("%.2f mol", float32(d))
}

func (d *DPT_14004) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14004(data)
		return nil
	case float64:
		*d = DPT_14004(data)
		return nil
	case int:
		*d = DPT_14004(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14004) Get() interface{} {
	return float32(d)
}

func (d DPT_14004) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14005 represents DPT 14.005 / Amplitude
type DPT_14005 float32

func (d DPT_14005) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14005) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14005(value)

	return nil
}

func (d DPT_14005) Unit() string {
	return ""
}

func (d DPT_14005) String() string {
	return fmt.Sprintf("%.2f", float32(d))
}

func (d *DPT_14005) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14005(data)
		return nil
	case float64:
		*d = DPT_14005(data)
		return nil
	case int:
		*d = DPT_14005(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14005) Get() interface{} {
	return float32(d)
}

func (d DPT_14005) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14006 represents DPT 14.006 / AngleRad
type DPT_14006 float32

func (d DPT_14006) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14006) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14006(value)

	return nil
}

func (d DPT_14006) Unit() string {
	return "rad"
}

func (d DPT_14006) String() string {
	return fmt.Sprintf("%.2f rad", float32(d))
}

func (d *DPT_14006) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14006(data)
		return nil
	case float64:
		*d = DPT_14006(data)
		return nil
	case int:
		*d = DPT_14006(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14006) Get() interface{} {
	return float32(d)
}

func (d DPT_14006) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14007 represents DPT 14.007 / AngleDeg
type DPT_14007 float32

func (d DPT_14007) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14007) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14007(value)

	return nil
}

func (d DPT_14007) Unit() string {
	return "°"
}

func (d DPT_14007) String() string {
	return fmt.Sprintf("%.2f °", float32(d))
}

func (d *DPT_14007) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14007(data)
		return nil
	case float64:
		*d = DPT_14007(data)
		return nil
	case int:
		*d = DPT_14007(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14007) Get() interface{} {
	return float32(d)
}

func (d DPT_14007) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14008 represents DPT 14.008 / Angular Momentum
type DPT_14008 float32

func (d DPT_14008) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14008) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14008(value)

	return nil
}

func (d DPT_14008) Unit() string {
	return "J s"
}

func (d DPT_14008) String() string {
	return fmt.Sprintf("%.2f J s", float32(d))
}

func (d *DPT_14008) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14008(data)
		return nil
	case float64:
		*d = DPT_14008(data)
		return nil
	case int:
		*d = DPT_14008(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14008) Get() interface{} {
	return float32(d)
}

func (d DPT_14008) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14009 represents DPT 14.009 / Angular Velocity
type DPT_14009 float32

func (d DPT_14009) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14009) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14009(value)

	return nil
}

func (d DPT_14009) Unit() string {
	return "rad/s"
}

func (d DPT_14009) String() string {
	return fmt.Sprintf("%.2f rad/s", float32(d))
}

func (d *DPT_14009) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14009(data)
		return nil
	case float64:
		*d = DPT_14009(data)
		return nil
	case int:
		*d = DPT_14009(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14009) Get() interface{} {
	return float32(d)
}

func (d DPT_14009) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14010 represents DPT 14.010 / Area
type DPT_14010 float32

func (d DPT_14010) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14010) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14010(value)

	return nil
}

func (d DPT_14010) Unit() string {
	return "m²"
}

func (d DPT_14010) String() string {
	return fmt.Sprintf("%.2f m²", float32(d))
}

func (d *DPT_14010) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14010(data)
		return nil
	case float64:
		*d = DPT_14010(data)
		return nil
	case int:
		*d = DPT_14010(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14010) Get() interface{} {
	return float32(d)
}

func (d DPT_14010) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14011 represents DPT 14.011 / Capacitance
type DPT_14011 float32

func (d DPT_14011) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14011) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14011(value)

	return nil
}

func (d DPT_14011) Unit() string {
	return "F"
}

func (d DPT_14011) String() string {
	return fmt.Sprintf("%.2f F", float32(d))
}

func (d *DPT_14011) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14011(data)
		return nil
	case float64:
		*d = DPT_14011(data)
		return nil
	case int:
		*d = DPT_14011(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14011) Get() interface{} {
	return float32(d)
}

func (d DPT_14011) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14012 represents DPT 14.012 / Charge DensitySurface
type DPT_14012 float32

func (d DPT_14012) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14012) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14012(value)

	return nil
}

func (d DPT_14012) Unit() string {
	return "C/m²"
}

func (d DPT_14012) String() string {
	return fmt.Sprintf("%.2f C/m²", float32(d))
}

func (d *DPT_14012) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14012(data)
		return nil
	case float64:
		*d = DPT_14012(data)
		return nil
	case int:
		*d = DPT_14012(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14012) Get() interface{} {
	return float32(d)
}

func (d DPT_14012) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14013 represents DPT 14.013 / Charge DensityVolume
type DPT_14013 float32

func (d DPT_14013) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14013) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14013(value)

	return nil
}

func (d DPT_14013) Unit() string {
	return "C/m³"
}

func (d DPT_14013) String() string {
	return fmt.Sprintf("%.2f C/m³", float32(d))
}

func (d *DPT_14013) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14013(data)
		return nil
	case float64:
		*d = DPT_14013(data)
		return nil
	case int:
		*d = DPT_14013(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14013) Get() interface{} {
	return float32(d)
}

func (d DPT_14013) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14014 represents DPT 14.014 / Compressibility
type DPT_14014 float32

func (d DPT_14014) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14014) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14014(value)

	return nil
}

func (d DPT_14014) Unit() string {
	return "m²/N"
}

func (d DPT_14014) String() string {
	return fmt.Sprintf("%.2f m²/N", float32(d))
}

func (d *DPT_14014) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14014(data)
		return nil
	case float64:
		*d = DPT_14014(data)
		return nil
	case int:
		*d = DPT_14014(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14014) Get() interface{} {
	return float32(d)
}

func (d DPT_14014) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14015 represents DPT 14.015 / Conductance
type DPT_14015 float32

func (d DPT_14015) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14015) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14015(value)

	return nil
}

func (d *DPT_14015) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14015(data)
		return nil
	case float64:
		*d = DPT_14015(data)
		return nil
	case int:
		*d = DPT_14015(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14015) Get() interface{} {
	return float32(d)
}

func (d DPT_14015) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

func (d DPT_14015) Unit() string {
	return "S"
}

func (d DPT_14015) String() string {
	return fmt.Sprintf("%.2f S", float32(d))
}

// DPT_14016 represents DPT 14.016 / Electrical Conductivity
type DPT_14016 float32

func (d DPT_14016) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14016) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14016(value)

	return nil
}

func (d DPT_14016) Unit() string {
	return "S/m"
}

func (d DPT_14016) String() string {
	return fmt.Sprintf("%.2f S/m", float32(d))
}

func (d *DPT_14016) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14016(data)
		return nil
	case float64:
		*d = DPT_14016(data)
		return nil
	case int:
		*d = DPT_14016(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14016) Get() interface{} {
	return float32(d)
}

func (d DPT_14016) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14017 represents DPT 14.017 / Density
type DPT_14017 float32

func (d DPT_14017) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14017) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14017(value)

	return nil
}

func (d DPT_14017) Unit() string {
	return "kg/m³"
}

func (d DPT_14017) String() string {
	return fmt.Sprintf("%.2f kg/m³", float32(d))
}

func (d *DPT_14017) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14017(data)
		return nil
	case float64:
		*d = DPT_14017(data)
		return nil
	case int:
		*d = DPT_14017(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14017) Get() interface{} {
	return float32(d)
}

func (d DPT_14017) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14018 represents DPT 14.018 / Electric Charge
type DPT_14018 float32

func (d DPT_14018) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14018) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14018(value)

	return nil
}

func (d DPT_14018) Unit() string {
	return "C"
}

func (d DPT_14018) String() string {
	return fmt.Sprintf("%.2f C", float32(d))
}

func (d *DPT_14018) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14018(data)
		return nil
	case float64:
		*d = DPT_14018(data)
		return nil
	case int:
		*d = DPT_14018(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14018) Get() interface{} {
	return float32(d)
}

func (d DPT_14018) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14019 represents DPT 14.019 / Electric Current
type DPT_14019 float32

func (d DPT_14019) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14019) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14019(value)

	return nil
}

func (d DPT_14019) Unit() string {
	return "A"
}

func (d DPT_14019) String() string {
	return fmt.Sprintf("%.2f A", float32(d))
}

func (d *DPT_14019) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14019(data)
		return nil
	case float64:
		*d = DPT_14019(data)
		return nil
	case int:
		*d = DPT_14019(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14019) Get() interface{} {
	return float32(d)
}

func (d DPT_14019) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14020 represents DPT 14.020 / Electric CurrentDensity
type DPT_14020 float32

func (d DPT_14020) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14020) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14020(value)

	return nil
}

func (d DPT_14020) Unit() string {
	return "A/m²"
}

func (d DPT_14020) String() string {
	return fmt.Sprintf("%.2f A/m²", float32(d))
}

func (d *DPT_14020) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14020(data)
		return nil
	case float64:
		*d = DPT_14020(data)
		return nil
	case int:
		*d = DPT_14020(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14020) Get() interface{} {
	return float32(d)
}

func (d DPT_14020) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14021 represents DPT 14.021 / Electric DipoleMoment
type DPT_14021 float32

func (d DPT_14021) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14021) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14021(value)

	return nil
}

func (d DPT_14021) Unit() string {
	return "C.m"
}

func (d DPT_14021) String() string {
	return fmt.Sprintf("%.2f C.m", float32(d))
}

func (d *DPT_14021) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14021(data)
		return nil
	case float64:
		*d = DPT_14021(data)
		return nil
	case int:
		*d = DPT_14021(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14021) Get() interface{} {
	return float32(d)
}

func (d DPT_14021) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14022 represents DPT 14.022 / Electric Displacement
type DPT_14022 float32

func (d DPT_14022) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14022) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14022(value)

	return nil
}

func (d DPT_14022) Unit() string {
	return "C/m²"
}

func (d DPT_14022) String() string {
	return fmt.Sprintf("%.2f C/m²", float32(d))
}

func (d *DPT_14022) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14022(data)
		return nil
	case float64:
		*d = DPT_14022(data)
		return nil
	case int:
		*d = DPT_14022(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14022) Get() interface{} {
	return float32(d)
}

func (d DPT_14022) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14023 represents DPT 14.023 / Electric FieldStrength
type DPT_14023 float32

func (d DPT_14023) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14023) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14023(value)

	return nil
}

func (d DPT_14023) Unit() string {
	return "V/m"
}

func (d DPT_14023) String() string {
	return fmt.Sprintf("%.2f V/m", float32(d))
}

func (d *DPT_14023) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14023(data)
		return nil
	case float64:
		*d = DPT_14023(data)
		return nil
	case int:
		*d = DPT_14023(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14023) Get() interface{} {
	return float32(d)
}

func (d DPT_14023) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14024 represents DPT 14.024 / Electric Flux
type DPT_14024 float32

func (d DPT_14024) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14024) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14024(value)

	return nil
}

func (d DPT_14024) Unit() string {
	return "c"
}

func (d DPT_14024) String() string {
	return fmt.Sprintf("%.2f c", float32(d))
}

func (d *DPT_14024) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14024(data)
		return nil
	case float64:
		*d = DPT_14024(data)
		return nil
	case int:
		*d = DPT_14024(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14024) Get() interface{} {
	return float32(d)
}

func (d DPT_14024) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14025 represents DPT 14.025 / Electric FluxDensity
type DPT_14025 float32

func (d DPT_14025) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14025) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14025(value)

	return nil
}

func (d DPT_14025) Unit() string {
	return "C/m²"
}

func (d DPT_14025) String() string {
	return fmt.Sprintf("%.2f C/m²", float32(d))
}

func (d *DPT_14025) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14025(data)
		return nil
	case float64:
		*d = DPT_14025(data)
		return nil
	case int:
		*d = DPT_14025(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14025) Get() interface{} {
	return float32(d)
}

func (d DPT_14025) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14026 represents DPT 14.026 / Electric Polarization
type DPT_14026 float32

func (d DPT_14026) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14026) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14026(value)

	return nil
}

func (d DPT_14026) Unit() string {
	return "C/m²"
}

func (d DPT_14026) String() string {
	return fmt.Sprintf("%.2f C/m²", float32(d))
}

func (d *DPT_14026) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14026(data)
		return nil
	case float64:
		*d = DPT_14026(data)
		return nil
	case int:
		*d = DPT_14026(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14026) Get() interface{} {
	return float32(d)
}

func (d DPT_14026) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14027 represents DPT 14.027 / Electric Potential
type DPT_14027 float32

func (d DPT_14027) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14027) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14027(value)

	return nil
}

func (d DPT_14027) Unit() string {
	return "V"
}

func (d DPT_14027) String() string {
	return fmt.Sprintf("%.2f V", float32(d))
}

func (d *DPT_14027) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14027(data)
		return nil
	case float64:
		*d = DPT_14027(data)
		return nil
	case int:
		*d = DPT_14027(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14027) Get() interface{} {
	return float32(d)
}

func (d DPT_14027) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14028 represents DPT 14.028 / Electric PotentialDifference
type DPT_14028 float32

func (d DPT_14028) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14028) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14028(value)

	return nil
}

func (d DPT_14028) Unit() string {
	return "V"
}

func (d DPT_14028) String() string {
	return fmt.Sprintf("%.2f V", float32(d))
}

func (d *DPT_14028) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14028(data)
		return nil
	case float64:
		*d = DPT_14028(data)
		return nil
	case int:
		*d = DPT_14028(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14028) Get() interface{} {
	return float32(d)
}

func (d DPT_14028) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14029 represents DPT 14.029 / ElectromagneticMoment
type DPT_14029 float32

func (d DPT_14029) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14029) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14029(value)

	return nil
}

func (d DPT_14029) Unit() string {
	return "A.m²"
}

func (d DPT_14029) String() string {
	return fmt.Sprintf("%.2f A.m²", float32(d))
}

func (d *DPT_14029) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14029(data)
		return nil
	case float64:
		*d = DPT_14029(data)
		return nil
	case int:
		*d = DPT_14029(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14029) Get() interface{} {
	return float32(d)
}

func (d DPT_14029) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14030 represents DPT 14.030 / Electromotive_Force
type DPT_14030 float32

func (d DPT_14030) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14030) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14030(value)

	return nil
}

func (d DPT_14030) Unit() string {
	return "V"
}

func (d DPT_14030) String() string {
	return fmt.Sprintf("%.2f V", float32(d))
}

func (d *DPT_14030) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14030(data)
		return nil
	case float64:
		*d = DPT_14030(data)
		return nil
	case int:
		*d = DPT_14030(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14030) Get() interface{} {
	return float32(d)
}

func (d DPT_14030) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14031 represents DPT 14.031 / Energy
type DPT_14031 float32

func (d DPT_14031) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14031) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14031(value)

	return nil
}

func (d DPT_14031) Unit() string {
	return "J"
}

func (d DPT_14031) String() string {
	return fmt.Sprintf("%.2f J", float32(d))
}

func (d *DPT_14031) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14031(data)
		return nil
	case float64:
		*d = DPT_14031(data)
		return nil
	case int:
		*d = DPT_14031(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14031) Get() interface{} {
	return float32(d)
}

func (d DPT_14031) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14032 represents DPT 14.032 / Force
type DPT_14032 float32

func (d DPT_14032) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14032) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14032(value)

	return nil
}

func (d DPT_14032) Unit() string {
	return "N"
}

func (d DPT_14032) String() string {
	return fmt.Sprintf("%.2f N", float32(d))
}

func (d *DPT_14032) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14032(data)
		return nil
	case float64:
		*d = DPT_14032(data)
		return nil
	case int:
		*d = DPT_14032(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14032) Get() interface{} {
	return float32(d)
}

func (d DPT_14032) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14033 represents DPT 14.033 / Frequency
type DPT_14033 float32

func (d DPT_14033) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14033) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14033(value)

	return nil
}

func (d DPT_14033) Unit() string {
	return "Hz"
}

func (d DPT_14033) String() string {
	return fmt.Sprintf("%.2f Hz", float32(d))
}

func (d *DPT_14033) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14033(data)
		return nil
	case float64:
		*d = DPT_14033(data)
		return nil
	case int:
		*d = DPT_14033(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14033) Get() interface{} {
	return float32(d)
}

func (d DPT_14033) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14034 represents DPT 14.034 / Angular Frequency
type DPT_14034 float32

func (d DPT_14034) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14034) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14034(value)

	return nil
}

func (d DPT_14034) Unit() string {
	return "rad/s"
}

func (d DPT_14034) String() string {
	return fmt.Sprintf("%.2f rad/s", float32(d))
}

func (d *DPT_14034) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14034(data)
		return nil
	case float64:
		*d = DPT_14034(data)
		return nil
	case int:
		*d = DPT_14034(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14034) Get() interface{} {
	return float32(d)
}

func (d DPT_14034) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14035 represents DPT 14.035 / Heat Capacity
type DPT_14035 float32

func (d DPT_14035) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14035) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14035(value)

	return nil
}

func (d DPT_14035) Unit() string {
	return "J/K"
}

func (d DPT_14035) String() string {
	return fmt.Sprintf("%.2f J/K", float32(d))
}

func (d *DPT_14035) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14035(data)
		return nil
	case float64:
		*d = DPT_14035(data)
		return nil
	case int:
		*d = DPT_14035(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14035) Get() interface{} {
	return float32(d)
}

func (d DPT_14035) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14036 represents DPT 14.036 / Heat Flow Rate
type DPT_14036 float32

func (d DPT_14036) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14036) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14036(value)

	return nil
}

func (d DPT_14036) Unit() string {
	return "W"
}

func (d DPT_14036) String() string {
	return fmt.Sprintf("%.2f W", float32(d))
}

func (d *DPT_14036) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14036(data)
		return nil
	case float64:
		*d = DPT_14036(data)
		return nil
	case int:
		*d = DPT_14036(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14036) Get() interface{} {
	return float32(d)
}

func (d DPT_14036) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14037 represents DPT 14.037 / Heat Quantity
type DPT_14037 float32

func (d DPT_14037) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14037) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14037(value)

	return nil
}

func (d DPT_14037) Unit() string {
	return "J"
}

func (d DPT_14037) String() string {
	return fmt.Sprintf("%.2f J", float32(d))
}

func (d *DPT_14037) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14037(data)
		return nil
	case float64:
		*d = DPT_14037(data)
		return nil
	case int:
		*d = DPT_14037(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14037) Get() interface{} {
	return float32(d)
}

func (d DPT_14037) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14038 represents DPT 14.038 / Impedance
type DPT_14038 float32

func (d DPT_14038) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14038) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14038(value)

	return nil
}

func (d DPT_14038) Unit() string {
	return "Ω"
}

func (d DPT_14038) String() string {
	return fmt.Sprintf("%.2f Ω", float32(d))
}

func (d *DPT_14038) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14038(data)
		return nil
	case float64:
		*d = DPT_14038(data)
		return nil
	case int:
		*d = DPT_14038(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14038) Get() interface{} {
	return float32(d)
}

func (d DPT_14038) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14039 represents DPT 14.039 / Length
type DPT_14039 float32

func (d DPT_14039) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14039) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14039(value)

	return nil
}

func (d DPT_14039) Unit() string {
	return "m"
}

func (d DPT_14039) String() string {
	return fmt.Sprintf("%.2f m", float32(d))
}

func (d *DPT_14039) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14039(data)
		return nil
	case float64:
		*d = DPT_14039(data)
		return nil
	case int:
		*d = DPT_14039(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14039) Get() interface{} {
	return float32(d)
}

func (d DPT_14039) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14040 represents DPT 14.040 / Light_Quantity
type DPT_14040 float32

func (d DPT_14040) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14040) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14040(value)

	return nil
}

func (d DPT_14040) Unit() string {
	return "lm.s"
}

func (d DPT_14040) String() string {
	return fmt.Sprintf("%.2f lm.s", float32(d))
}

func (d *DPT_14040) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14040(data)
		return nil
	case float64:
		*d = DPT_14040(data)
		return nil
	case int:
		*d = DPT_14040(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14040) Get() interface{} {
	return float32(d)
}

func (d DPT_14040) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14041 represents DPT 14.041 / Luminance
type DPT_14041 float32

func (d DPT_14041) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14041) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14041(value)

	return nil
}

func (d DPT_14041) Unit() string {
	return "cd/m²"
}

func (d DPT_14041) String() string {
	return fmt.Sprintf("%.2f cd/m²", float32(d))
}

func (d *DPT_14041) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14041(data)
		return nil
	case float64:
		*d = DPT_14041(data)
		return nil
	case int:
		*d = DPT_14041(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14041) Get() interface{} {
	return float32(d)
}

func (d DPT_14041) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14042 represents DPT 14.042 / Luminous Flux
type DPT_14042 float32

func (d DPT_14042) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14042) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14042(value)

	return nil
}

func (d DPT_14042) Unit() string {
	return "lm"
}

func (d DPT_14042) String() string {
	return fmt.Sprintf("%.2f lm", float32(d))
}

func (d *DPT_14042) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14042(data)
		return nil
	case float64:
		*d = DPT_14042(data)
		return nil
	case int:
		*d = DPT_14042(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14042) Get() interface{} {
	return float32(d)
}

func (d DPT_14042) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14043 represents DPT 14.043 / Luminous Intensity
type DPT_14043 float32

func (d DPT_14043) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14043) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14043(value)

	return nil
}

func (d DPT_14043) Unit() string {
	return "cd"
}

func (d DPT_14043) String() string {
	return fmt.Sprintf("%.2f cd", float32(d))
}

func (d *DPT_14043) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14043(data)
		return nil
	case float64:
		*d = DPT_14043(data)
		return nil
	case int:
		*d = DPT_14043(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14043) Get() interface{} {
	return float32(d)
}

func (d DPT_14043) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14044 represents DPT 14.044 / Magnetic FieldStrength
type DPT_14044 float32

func (d DPT_14044) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14044) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14044(value)

	return nil
}

func (d DPT_14044) Unit() string {
	return "A/m"
}

func (d DPT_14044) String() string {
	return fmt.Sprintf("%.2f A/m", float32(d))
}

func (d *DPT_14044) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14044(data)
		return nil
	case float64:
		*d = DPT_14044(data)
		return nil
	case int:
		*d = DPT_14044(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14044) Get() interface{} {
	return float32(d)
}

func (d DPT_14044) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14045 represents DPT 14.045 / Magnetic Flux
type DPT_14045 float32

func (d DPT_14045) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14045) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14045(value)

	return nil
}

func (d DPT_14045) Unit() string {
	return "Wb"
}

func (d DPT_14045) String() string {
	return fmt.Sprintf("%.2f Wb", float32(d))
}

func (d *DPT_14045) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14045(data)
		return nil
	case float64:
		*d = DPT_14045(data)
		return nil
	case int:
		*d = DPT_14045(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14045) Get() interface{} {
	return float32(d)
}

func (d DPT_14045) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14046 represents DPT 14.046 / Magnetic FluxDensity
type DPT_14046 float32

func (d DPT_14046) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14046) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14046(value)

	return nil
}

func (d DPT_14046) Unit() string {
	return "T"
}

func (d DPT_14046) String() string {
	return fmt.Sprintf("%.2f T", float32(d))
}

func (d *DPT_14046) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14046(data)
		return nil
	case float64:
		*d = DPT_14046(data)
		return nil
	case int:
		*d = DPT_14046(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14046) Get() interface{} {
	return float32(d)
}

func (d DPT_14046) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14047 represents DPT 14.047 / Magnetic Moment
type DPT_14047 float32

func (d DPT_14047) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14047) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14047(value)

	return nil
}

func (d DPT_14047) Unit() string {
	return "A.m²"
}

func (d DPT_14047) String() string {
	return fmt.Sprintf("%.2f A.m²", float32(d))
}

func (d *DPT_14047) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14047(data)
		return nil
	case float64:
		*d = DPT_14047(data)
		return nil
	case int:
		*d = DPT_14047(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14047) Get() interface{} {
	return float32(d)
}

func (d DPT_14047) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14048 represents DPT 14.048 / Magnetic Polarization
type DPT_14048 float32

func (d DPT_14048) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14048) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14048(value)

	return nil
}

func (d DPT_14048) Unit() string {
	return "T"
}

func (d DPT_14048) String() string {
	return fmt.Sprintf("%.2f T", float32(d))
}

func (d *DPT_14048) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14048(data)
		return nil
	case float64:
		*d = DPT_14048(data)
		return nil
	case int:
		*d = DPT_14048(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14048) Get() interface{} {
	return float32(d)
}

func (d DPT_14048) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14049 represents DPT 14.049 / Magnetization
type DPT_14049 float32

func (d DPT_14049) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14049) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14049(value)

	return nil
}

func (d DPT_14049) Unit() string {
	return "A/m"
}

func (d DPT_14049) String() string {
	return fmt.Sprintf("%.2f A/m", float32(d))
}

func (d *DPT_14049) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14049(data)
		return nil
	case float64:
		*d = DPT_14049(data)
		return nil
	case int:
		*d = DPT_14049(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14049) Get() interface{} {
	return float32(d)
}

func (d DPT_14049) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14050 represents DPT 14.050 / MagnetomotiveForce
type DPT_14050 float32

func (d DPT_14050) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14050) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14050(value)

	return nil
}

func (d DPT_14050) Unit() string {
	return "A"
}

func (d DPT_14050) String() string {
	return fmt.Sprintf("%.2f A", float32(d))
}

func (d *DPT_14050) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14050(data)
		return nil
	case float64:
		*d = DPT_14050(data)
		return nil
	case int:
		*d = DPT_14050(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14050) Get() interface{} {
	return float32(d)
}

func (d DPT_14050) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14051 represents DPT 14.051 / Mass
type DPT_14051 float32

func (d DPT_14051) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14051) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14051(value)

	return nil
}

func (d DPT_14051) Unit() string {
	return "kg"
}

func (d DPT_14051) String() string {
	return fmt.Sprintf("%.2f kg", float32(d))
}

func (d *DPT_14051) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14051(data)
		return nil
	case float64:
		*d = DPT_14051(data)
		return nil
	case int:
		*d = DPT_14051(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14051) Get() interface{} {
	return float32(d)
}

func (d DPT_14051) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14052 represents DPT 14.052 / MassFlux
type DPT_14052 float32

func (d DPT_14052) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14052) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14052(value)

	return nil
}

func (d DPT_14052) Unit() string {
	return "kg/s"
}

func (d DPT_14052) String() string {
	return fmt.Sprintf("%.2f kg/s", float32(d))
}

func (d *DPT_14052) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14052(data)
		return nil
	case float64:
		*d = DPT_14052(data)
		return nil
	case int:
		*d = DPT_14052(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14052) Get() interface{} {
	return float32(d)
}

func (d DPT_14052) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14053 represents DPT 14.053 / Momentum
type DPT_14053 float32

func (d DPT_14053) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14053) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14053(value)

	return nil
}

func (d DPT_14053) Unit() string {
	return "N/s"
}

func (d DPT_14053) String() string {
	return fmt.Sprintf("%.2f N/s", float32(d))
}

func (d *DPT_14053) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14053(data)
		return nil
	case float64:
		*d = DPT_14053(data)
		return nil
	case int:
		*d = DPT_14053(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14053) Get() interface{} {
	return float32(d)
}

func (d DPT_14053) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14054 represents DPT 14.054 / Phase Angle, Radiant
type DPT_14054 float32

func (d DPT_14054) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14054) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14054(value)

	return nil
}

func (d DPT_14054) Unit() string {
	return "rad"
}

func (d DPT_14054) String() string {
	return fmt.Sprintf("%.2f rad", float32(d))
}

func (d *DPT_14054) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14054(data)
		return nil
	case float64:
		*d = DPT_14054(data)
		return nil
	case int:
		*d = DPT_14054(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14054) Get() interface{} {
	return float32(d)
}

func (d DPT_14054) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14055 represents DPT 14.055 / Phase Angle, Degree
type DPT_14055 float32

func (d DPT_14055) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14055) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14055(value)

	return nil
}

func (d DPT_14055) Unit() string {
	return "°"
}

func (d DPT_14055) String() string {
	return fmt.Sprintf("%.2f °", float32(d))
}

func (d *DPT_14055) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14055(data)
		return nil
	case float64:
		*d = DPT_14055(data)
		return nil
	case int:
		*d = DPT_14055(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14055) Get() interface{} {
	return float32(d)
}

func (d DPT_14055) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14056 represents DPT 14.056 / Power
type DPT_14056 float32

func (d DPT_14056) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14056) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14056(value)

	return nil
}

func (d DPT_14056) Unit() string {
	return "W"
}

func (d DPT_14056) String() string {
	return fmt.Sprintf("%.2f W", float32(d))
}

func (d *DPT_14056) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14056(data)
		return nil
	case float64:
		*d = DPT_14056(data)
		return nil
	case int:
		*d = DPT_14056(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14056) Get() interface{} {
	return float32(d)
}

func (d DPT_14056) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14057 represents DPT 14.057 / Power Factor
type DPT_14057 float32

func (d DPT_14057) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14057) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14057(value)

	return nil
}

func (d DPT_14057) Unit() string {
	return "cosΦ"
}

func (d DPT_14057) String() string {
	return fmt.Sprintf("%.2f cosΦ", float32(d))
}

func (d *DPT_14057) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14057(data)
		return nil
	case float64:
		*d = DPT_14057(data)
		return nil
	case int:
		*d = DPT_14057(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14057) Get() interface{} {
	return float32(d)
}

func (d DPT_14057) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14058 represents DPT 14.058 / Pressure
type DPT_14058 float32

func (d DPT_14058) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14058) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14058(value)

	return nil
}

func (d DPT_14058) Unit() string {
	return "Pa"
}

func (d DPT_14058) String() string {
	return fmt.Sprintf("%.2f Pa", float32(d))
}

func (d *DPT_14058) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14058(data)
		return nil
	case float64:
		*d = DPT_14058(data)
		return nil
	case int:
		*d = DPT_14058(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14058) Get() interface{} {
	return float32(d)
}

func (d DPT_14058) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14059 represents DPT 14.059 / Reactance
type DPT_14059 float32

func (d DPT_14059) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14059) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14059(value)

	return nil
}

func (d DPT_14059) Unit() string {
	return "Ω"
}

func (d DPT_14059) String() string {
	return fmt.Sprintf("%.2f Ω", float32(d))
}

func (d *DPT_14059) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14059(data)
		return nil
	case float64:
		*d = DPT_14059(data)
		return nil
	case int:
		*d = DPT_14059(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14059) Get() interface{} {
	return float32(d)
}

func (d DPT_14059) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14060 represents DPT 14.060 / Resistance
type DPT_14060 float32

func (d DPT_14060) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14060) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14060(value)

	return nil
}

func (d DPT_14060) Unit() string {
	return "Ω"
}

func (d DPT_14060) String() string {
	return fmt.Sprintf("%.2f Ω", float32(d))
}

func (d *DPT_14060) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14060(data)
		return nil
	case float64:
		*d = DPT_14060(data)
		return nil
	case int:
		*d = DPT_14060(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14060) Get() interface{} {
	return float32(d)
}

func (d DPT_14060) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14061 represents DPT 14.061 / Resistivity
type DPT_14061 float32

func (d DPT_14061) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14061) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14061(value)

	return nil
}

func (d DPT_14061) Unit() string {
	return "Ω.m"
}

func (d DPT_14061) String() string {
	return fmt.Sprintf("%.2f Ω.m", float32(d))
}

func (d *DPT_14061) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14061(data)
		return nil
	case float64:
		*d = DPT_14061(data)
		return nil
	case int:
		*d = DPT_14061(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14061) Get() interface{} {
	return float32(d)
}

func (d DPT_14061) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14062 represents DPT 14.062 / SelfInductance
type DPT_14062 float32

func (d DPT_14062) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14062) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14062(value)

	return nil
}

func (d DPT_14062) Unit() string {
	return "H"
}

func (d DPT_14062) String() string {
	return fmt.Sprintf("%.2f H", float32(d))
}

func (d *DPT_14062) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14062(data)
		return nil
	case float64:
		*d = DPT_14062(data)
		return nil
	case int:
		*d = DPT_14062(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14062) Get() interface{} {
	return float32(d)
}

func (d DPT_14062) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14063 represents DPT 14.063 / SolidAngle
type DPT_14063 float32

func (d DPT_14063) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14063) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14063(value)

	return nil
}

func (d DPT_14063) Unit() string {
	return "sr"
}

func (d DPT_14063) String() string {
	return fmt.Sprintf("%.2f sr", float32(d))
}

func (d *DPT_14063) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14063(data)
		return nil
	case float64:
		*d = DPT_14063(data)
		return nil
	case int:
		*d = DPT_14063(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14063) Get() interface{} {
	return float32(d)
}

func (d DPT_14063) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14064 represents DPT 14.064 / Sound Intensity
type DPT_14064 float32

func (d DPT_14064) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14064) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14064(value)

	return nil
}

func (d DPT_14064) Unit() string {
	return "W/m²"
}

func (d DPT_14064) String() string {
	return fmt.Sprintf("%.2f W/m²", float32(d))
}

func (d *DPT_14064) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14064(data)
		return nil
	case float64:
		*d = DPT_14064(data)
		return nil
	case int:
		*d = DPT_14064(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14064) Get() interface{} {
	return float32(d)
}

func (d DPT_14064) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14065 represents DPT 14.065 / Speed
type DPT_14065 float32

func (d DPT_14065) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14065) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14065(value)

	return nil
}

func (d DPT_14065) Unit() string {
	return "m/s"
}

func (d DPT_14065) String() string {
	return fmt.Sprintf("%.2f m/s", float32(d))
}

func (d *DPT_14065) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14065(data)
		return nil
	case float64:
		*d = DPT_14065(data)
		return nil
	case int:
		*d = DPT_14065(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14065) Get() interface{} {
	return float32(d)
}

func (d DPT_14065) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14066 represents DPT 14.066 / Stress
type DPT_14066 float32

func (d DPT_14066) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14066) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14066(value)

	return nil
}

func (d DPT_14066) Unit() string {
	return "Pa"
}

func (d DPT_14066) String() string {
	return fmt.Sprintf("%.2f Pa", float32(d))
}

func (d *DPT_14066) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14066(data)
		return nil
	case float64:
		*d = DPT_14066(data)
		return nil
	case int:
		*d = DPT_14066(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14066) Get() interface{} {
	return float32(d)
}

func (d DPT_14066) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14067 represents DPT 14.067 / Surface Tension
type DPT_14067 float32

func (d DPT_14067) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14067) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14067(value)

	return nil
}

func (d DPT_14067) Unit() string {
	return "N/m"
}

func (d DPT_14067) String() string {
	return fmt.Sprintf("%.2f N/m", float32(d))
}

func (d *DPT_14067) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14067(data)
		return nil
	case float64:
		*d = DPT_14067(data)
		return nil
	case int:
		*d = DPT_14067(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14067) Get() interface{} {
	return float32(d)
}

func (d DPT_14067) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14068 represents DPT 14.068 / Common Temperature
type DPT_14068 float32

func (d DPT_14068) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14068) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14068(value)

	return nil
}

func (d DPT_14068) Unit() string {
	return "°C"
}

func (d DPT_14068) String() string {
	return fmt.Sprintf("%.2f °C", float32(d))
}

func (d *DPT_14068) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14068(data)
		return nil
	case float64:
		*d = DPT_14068(data)
		return nil
	case int:
		*d = DPT_14068(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14068) Get() interface{} {
	return float32(d)
}

func (d DPT_14068) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14069 represents DPT 14.069 / Absolute Temperature
type DPT_14069 float32

func (d DPT_14069) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14069) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14069(value)

	return nil
}

func (d DPT_14069) Unit() string {
	return "K"
}

func (d DPT_14069) String() string {
	return fmt.Sprintf("%.2f K", float32(d))
}

func (d *DPT_14069) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14069(data)
		return nil
	case float64:
		*d = DPT_14069(data)
		return nil
	case int:
		*d = DPT_14069(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14069) Get() interface{} {
	return float32(d)
}

func (d DPT_14069) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14070 represents DPT 14.070 / Temperature Difference
type DPT_14070 float32

func (d DPT_14070) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14070) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14070(value)

	return nil
}

func (d DPT_14070) Unit() string {
	return "K"
}

func (d DPT_14070) String() string {
	return fmt.Sprintf("%.2f K", float32(d))
}

func (d *DPT_14070) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14070(data)
		return nil
	case float64:
		*d = DPT_14070(data)
		return nil
	case int:
		*d = DPT_14070(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14070) Get() interface{} {
	return float32(d)
}

func (d DPT_14070) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14071 represents DPT 14.071 / Thermal Capacity
type DPT_14071 float32

func (d DPT_14071) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14071) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14071(value)

	return nil
}

func (d DPT_14071) Unit() string {
	return "J/K"
}

func (d DPT_14071) String() string {
	return fmt.Sprintf("%.2f J/K", float32(d))
}

func (d *DPT_14071) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14071(data)
		return nil
	case float64:
		*d = DPT_14071(data)
		return nil
	case int:
		*d = DPT_14071(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14071) Get() interface{} {
	return float32(d)
}

func (d DPT_14071) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14072 represents DPT 14.072 / Thermal Conductivity
type DPT_14072 float32

func (d DPT_14072) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14072) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14072(value)

	return nil
}

func (d DPT_14072) Unit() string {
	return "W/mK"
}

func (d DPT_14072) String() string {
	return fmt.Sprintf("%.2f W/mK", float32(d))
}

func (d *DPT_14072) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14072(data)
		return nil
	case float64:
		*d = DPT_14072(data)
		return nil
	case int:
		*d = DPT_14072(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14072) Get() interface{} {
	return float32(d)
}

func (d DPT_14072) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14073 represents DPT 14.073 / Thermoelectric Power
type DPT_14073 float32

func (d DPT_14073) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14073) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14073(value)

	return nil
}

func (d DPT_14073) Unit() string {
	return "V/K"
}

func (d DPT_14073) String() string {
	return fmt.Sprintf("%.2f V/K", float32(d))
}

func (d *DPT_14073) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14073(data)
		return nil
	case float64:
		*d = DPT_14073(data)
		return nil
	case int:
		*d = DPT_14073(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14073) Get() interface{} {
	return float32(d)
}

func (d DPT_14073) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14074 represents DPT 14.074 / Time
type DPT_14074 float32

func (d DPT_14074) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14074) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14074(value)

	return nil
}

func (d DPT_14074) Unit() string {
	return "s"
}

func (d DPT_14074) String() string {
	return fmt.Sprintf("%.2f s", float32(d))
}

func (d *DPT_14074) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14074(data)
		return nil
	case float64:
		*d = DPT_14074(data)
		return nil
	case int:
		*d = DPT_14074(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14074) Get() interface{} {
	return float32(d)
}

func (d DPT_14074) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14075 represents DPT 14.075 / Torque
type DPT_14075 float32

func (d DPT_14075) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14075) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14075(value)

	return nil
}

func (d DPT_14075) Unit() string {
	return "N.m"
}

func (d DPT_14075) String() string {
	return fmt.Sprintf("%.2f N.m", float32(d))
}

func (d *DPT_14075) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14075(data)
		return nil
	case float64:
		*d = DPT_14075(data)
		return nil
	case int:
		*d = DPT_14075(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14075) Get() interface{} {
	return float32(d)
}

func (d DPT_14075) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14076 represents DPT 14.076 / Volume
type DPT_14076 float32

func (d DPT_14076) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14076) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14076(value)

	return nil
}

func (d DPT_14076) Unit() string {
	return "m³"
}

func (d DPT_14076) String() string {
	return fmt.Sprintf("%.2f m³", float32(d))
}

func (d *DPT_14076) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14076(data)
		return nil
	case float64:
		*d = DPT_14076(data)
		return nil
	case int:
		*d = DPT_14076(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14076) Get() interface{} {
	return float32(d)
}

func (d DPT_14076) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14077 represents DPT 14.077 / Volume Flux
type DPT_14077 float32

func (d DPT_14077) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14077) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14077(value)

	return nil
}

func (d DPT_14077) Unit() string {
	return "m³/s"
}

func (d DPT_14077) String() string {
	return fmt.Sprintf("%.2f m³/s", float32(d))
}

func (d *DPT_14077) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14077(data)
		return nil
	case float64:
		*d = DPT_14077(data)
		return nil
	case int:
		*d = DPT_14077(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14077) Get() interface{} {
	return float32(d)
}

func (d DPT_14077) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14078 represents DPT 14.078 / Weight
type DPT_14078 float32

func (d DPT_14078) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14078) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14078(value)

	return nil
}

func (d DPT_14078) Unit() string {
	return "N"
}

func (d DPT_14078) String() string {
	return fmt.Sprintf("%.2f N", float32(d))
}

func (d *DPT_14078) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14078(data)
		return nil
	case float64:
		*d = DPT_14078(data)
		return nil
	case int:
		*d = DPT_14078(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14078) Get() interface{} {
	return float32(d)
}

func (d DPT_14078) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}

// DPT_14079 represents DPT 14.079 / Work
type DPT_14079 float32

func (d DPT_14079) Pack() []byte {
	return packF32(float32(d))
}

func (d *DPT_14079) Unpack(data []byte) error {
	var value float32

	if err := unpackF32(data, &value); err != nil {
		return err
	}

	*d = DPT_14079(value)

	return nil
}

func (d DPT_14079) Unit() string {
	return "J"
}

func (d DPT_14079) String() string {
	return fmt.Sprintf("%.2f J", float32(d))
}

func (d *DPT_14079) Set(v interface{}) error {
	switch data := v.(type) {
	case float32:
		*d = DPT_14079(data)
		return nil
	case float64:
		*d = DPT_14079(data)
		return nil
	case int:
		*d = DPT_14079(data)
		return nil
	}
	return fmt.Errorf("invalid value for %[1]T, %[2]T=%[2]v ", *d, v)
}

func (d DPT_14079) Get() interface{} {
	return float32(d)
}

func (d DPT_14079) Formatting(format string) string {
	if format != "" {
		return fmt.Sprintf(format, float32(d))
	}
	return fmt.Sprintf("%.2f", float32(d))
}
