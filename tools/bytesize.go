/*
		Convert byte size to bytes, kilobytes, megabytes, gigabytes, ...

		These utility functions convert a byte size to a human-readable string
		in either SI (decimal) or IEC (binary) format.


		Source: https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/


		These utility functions convert a size in bytes to a human-readable string in either SI (decimal) or IEC (binary) format.

			Input	ByteCountSI	ByteCountIEC
			999		"999 B"		"999 B"
			1000	"1.0 kB"	"1000 B"
			1023	"1.0 kB"	"1023 B"
			1024	"1.0 kB"	"1.0 KiB"
	 987,654,321	"987.7 MB"	"941.9 MiB"

math.MaxInt64	"9.2 EB"	"8.0 EiB"
*/
package tools

import (
	"fmt"
)

func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
