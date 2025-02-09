//go:build !windows

package shared

import "github.com/ebitengine/purego"

func Load(name string) (uintptr, error) {
	return purego.Dlopen(name, purego.RTLD_LAZY)
}
