package stdlog

type TimeFormat int

const (
	// TimeFormatUnix defines a time format that makes time fields to be
	// serialized as Unix timestamp integers.
	TimeFormatUnix TimeFormat = iota

	// TimeFormatUnixMs defines a time format that makes time fields to be
	// serialized as Unix timestamp integers in milliseconds.
	//TimeFormatUnixMs = "UNIXMS"
	TimeFormatUnixMs

	// TimeFormatUnixMicro defines a time format that makes time fields to be
	// serialized as Unix timestamp integers in microseconds.
	//TimeFormatUnixMicro = "UNIXMICRO"
	TimeFormatUnixMicro

	// TimeFormatUnixNano defines a time format that makes time fields to be
	// serialized as Unix timestamp integers in nanoseconds.
	//TimeFormatUnixNano = "UNIXNANO"
	TimeFormatUnixNano
)
