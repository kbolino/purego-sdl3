package convert

import (
	"strings"
	"syscall"
	"unsafe"
)

// ToBytePtr converts a Go string into a null-terminated C-style string.
// It returns an [syscall.EINVAL] error, if s contains a null byte at any location.
func ToBytePtr(s string) (*byte, error) {
	if strings.IndexByte(s, 0) != -1 {
		return nil, syscall.EINVAL
	}
	result := make([]byte, len(s)+1)
	copy(result, s)
	return &result[0], nil
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
