//go:build windows

package shared

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func Load(name string) (uintptr, error) {
	handle, err := windows.LoadLibrary(name)
	if err != nil {
		err = fmt.Errorf("%s: error loading library: %w", name, err)
	}
	return uintptr(handle), err
}
