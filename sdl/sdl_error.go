package sdl

import "fmt"

type Error struct {
	msg string
}

func (e Error) Error() string {
	return e.msg
}

// GetError retrieves a message about the last error that occurred on the current thread.
func GetError() Error {
	return Error{sdlGetError()}
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

func CheckBool(b bool) error {
	if !b {
		return GetError()
	} else {
		return nil
	}
}

func CheckPtr[T any](p *T) (*T, error) {
	if p == nil {
		return nil, GetError()
	} else {
		return p, nil
	}
}
