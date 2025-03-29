package sdl

import "fmt"

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

// SetError sets the SDL error message for the current thread.
func SetError(format string, a ...any) bool {
	return sdlSetError(fmt.Sprintf(format, a...))
}

// func SetErrorV(fmt string, ap va_list) bool {
//	return sdlSetErrorV(fmt, ap)
// }

func InvalidParamError(param string) bool {
	return SetError("Parameter '%s' is invalid", param)
}

type Err struct {
	msg string
}

func (e Err) Error() string {
	return "SDL error: " + e.msg
}

func checkBool(b bool) error {
	if !b {
		return getErr()
	}
	return nil
}

func checkPtr[T any](p *T) (*T, error) {
	if p == nil {
		return nil, getErr()
	} else {
		return p, nil
	}
}

func getErr() Err {
	return Err{msg: GetError()}
}
