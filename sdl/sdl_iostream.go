package sdl

import "unsafe"

type IOStream unsafe.Pointer

// IOFromConstMem returns a read-only memory buffer for use with [IOStream] or nil on failure.
func IOFromConstMem(mem []byte) IOStream {
	return sdlIOFromConstMem(mem, len(mem))
}

// CloseIO closes and frees an allocated [IOStream] structure.
func CloseIO(context IOStream) bool {
	return sdlCloseIO(context)
}
