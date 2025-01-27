//go:build unix

package convert

import "golang.org/x/sys/unix"

func ToBytePtr(s string) (*byte, error) {
	return unix.BytePtrFromString(s)
}

func ToString(p *byte) string {
	return unix.BytePtrToString(p)
}
