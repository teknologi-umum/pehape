package pehape

import (
	"crypto/md5"
	"encoding/hex"
)

// The Md5() function calculates the MD5 hash of a string.
// Parameter
//   - str => string to be calculated
//   - raw => (optional). Specifies hex or binary output format:
//     TRUE: Raw 16 character binary format
//     FALSE: Default. 32 character hex number
//
// Return
// - res => calculated MD5 hash
func Md5(str string, raw ...bool) (res string) {
	var defaultRaw bool
	// if raw parameter more than 1, only get the first
	if len(raw) > 0 {
		defaultRaw = raw[0]
	}

	hash := md5.Sum([]byte(str))
	if !defaultRaw {
		return hex.EncodeToString(hash[:])
	}
	return string(hash[:])
}
