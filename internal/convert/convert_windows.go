package convert

import "golang.org/x/sys/windows"

func ToBytePtr(s string) (*byte, error) {
	return windows.BytePtrFromString(s)
}

func ToString(p *byte) string {
	return windows.BytePtrToString(p)
}
