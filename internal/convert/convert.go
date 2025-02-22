package convert

import (
	"strings"
	"unsafe"
)

// ToBytePtr converts a Go string to a null-terminated C-style string by just appending a null byte,
// if s doesn't already contain one.
func ToBytePtr(s string) *byte {
	size := len(s) + 1
	if index := strings.IndexByte(s, 0); index != -1 {
		size = index + 1
	}

	result := make([]byte, size)
	copy(result, s)
	return &result[0]
}

// ToString converts a null-terminated C-style string into a Go string.
func ToString(p *byte) string {
	if p == nil {
		return ""
	}
	i := 0
	for ptr := unsafe.Pointer(p); *(*byte)(unsafe.Add(ptr, i)) != 0; i++ {
	}
	return string(unsafe.Slice(p, i))
}
