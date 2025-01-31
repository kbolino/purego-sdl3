package sdl

// GetError retrieves a message about the last error that occurred on the current thread.
func GetError() string {
	return sdlGetError()
}

// ClearError clears any previous error message for this thread.
func ClearError() bool {
	return sdlClearError()
}

// func OutOfMemory() bool {
//	return sdlOutOfMemory()
// }

// func SetError(fmt string) bool {
//	return sdlSetError(fmt)
// }

// func SetErrorV(fmt string, ap va_list) bool {
//	return sdlSetErrorV(fmt, ap)
// }
