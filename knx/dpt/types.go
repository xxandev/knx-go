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

func GetTypes() []DatapointValue {
	return []DatapointValue{
		// 1.xxx
		new(DPT_1001),
		new(DPT_1002),
		new(DPT_1003),
		new(DPT_1004),
		new(DPT_1005),
		new(DPT_1006),
		new(DPT_1007),
		new(DPT_1008),
		new(DPT_1009),
		new(DPT_1010),
		new(DPT_1011),
		new(DPT_1012),
		new(DPT_1013),
		new(DPT_1014),
		new(DPT_1015),
		new(DPT_1016),
		new(DPT_1017),
		new(DPT_1018),
		new(DPT_1019),
		new(DPT_1021),
		new(DPT_1022),
		new(DPT_1023),
		new(DPT_1024),
		new(DPT_1100),

		// 5.xxx
		new(DPT_5001),
		new(DPT_5003),
		new(DPT_5004),
		new(DPT_5005),

		// 6.xxx
		new(DPT_6010),

		// 7.xxx
		new(DPT_7001),
		new(DPT_7002),
		new(DPT_7003),
		new(DPT_7004),
		new(DPT_7005),
		new(DPT_7006),
		new(DPT_7007),
		new(DPT_7010),
		new(DPT_7011),
		new(DPT_7012),
		new(DPT_7013),

		// 9.xxx
		new(DPT_9001),
		new(DPT_9002),
		new(DPT_9003),
		new(DPT_9004),
		new(DPT_9005),
		new(DPT_9006),
		new(DPT_9007),
		new(DPT_9008),
		new(DPT_9010),
		new(DPT_9011),
		new(DPT_9020),
		new(DPT_9021),
		new(DPT_9022),
		new(DPT_9023),
		new(DPT_9024),
		new(DPT_9025),
		new(DPT_9026),
		new(DPT_9027),
		new(DPT_9028),

		// 10.xxx
		new(DPT_10001),

		// 11.xxx
		new(DPT_11001),

		// 12.xxx
		new(DPT_12001),

		// 13.xxx
		new(DPT_13001),
		new(DPT_13002),
		new(DPT_13010),
		new(DPT_13011),
		new(DPT_13012),
		new(DPT_13013),
		new(DPT_13014),
		new(DPT_13015),
		new(DPT_13016),
		new(DPT_13100),

		// 14.xxx
		new(DPT_14000),
		new(DPT_14001),
		new(DPT_14002),
		new(DPT_14003),
		new(DPT_14004),
		new(DPT_14005),
		new(DPT_14006),
		new(DPT_14007),
		new(DPT_14008),
		new(DPT_14009),
		new(DPT_14010),
		new(DPT_14011),
		new(DPT_14012),
		new(DPT_14013),
		new(DPT_14014),
		new(DPT_14015),
		new(DPT_14016),
		new(DPT_14017),
		new(DPT_14018),
		new(DPT_14019),
		new(DPT_14020),
		new(DPT_14021),
		new(DPT_14022),
		new(DPT_14023),
		new(DPT_14024),
		new(DPT_14025),
		new(DPT_14026),
		new(DPT_14027),
		new(DPT_14028),
		new(DPT_14029),
		new(DPT_14030),
		new(DPT_14031),
		new(DPT_14032),
		new(DPT_14033),
		new(DPT_14034),
		new(DPT_14035),
		new(DPT_14036),
		new(DPT_14037),
		new(DPT_14038),
		new(DPT_14039),
		new(DPT_14040),
		new(DPT_14041),
		new(DPT_14042),
		new(DPT_14043),
		new(DPT_14044),
		new(DPT_14045),
		new(DPT_14046),
		new(DPT_14047),
		new(DPT_14048),
		new(DPT_14049),
		new(DPT_14050),
		new(DPT_14051),
		new(DPT_14052),
		new(DPT_14053),
		new(DPT_14054),
		new(DPT_14055),
		new(DPT_14056),
		new(DPT_14057),
		new(DPT_14058),
		new(DPT_14059),
		new(DPT_14060),
		new(DPT_14061),
		new(DPT_14062),
		new(DPT_14063),
		new(DPT_14064),
		new(DPT_14065),
		new(DPT_14066),
		new(DPT_14067),
		new(DPT_14068),
		new(DPT_14069),
		new(DPT_14070),
		new(DPT_14071),
		new(DPT_14072),
		new(DPT_14073),
		new(DPT_14074),
		new(DPT_14075),
		new(DPT_14076),
		new(DPT_14077),
		new(DPT_14078),
		new(DPT_14079),

		// 16.xxx
		new(DPT_16000),
		new(DPT_16001),

		// 17.xxx
		new(DPT_17001),

		// 18.xxx
		new(DPT_18001),

		// 20.xxx
		new(DPT_20102),
		new(DPT_20105),

		// 28.xxx
		new(DPT_28001),

		// 242.xxx
		new(DPT_242600),

		// 251.xxx
		new(DPT_251600),
	}
}
