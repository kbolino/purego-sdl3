package sdl

import (
	"errors"
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/convert"
	"github.com/jupiterrider/purego-sdl3/internal/mem"
)

type MouseWheelDirection uint32

const (
	MouseWheelNormal MouseWheelDirection = iota
	MouseWheelFlipped
)

type SystemCursor uint32

const (
	SystemCursorDefault SystemCursor = iota
	SystemCursorText
	SystemCursorWait
	SystemCursorCrosshair
	SystemCursorProgress
	SystemCursorNwseResize
	SystemCursorNeswResize
	SystemCursorEwResize
	SystemCursorNsResize
	SystemCursorMove
	SystemCursorNotAllowed
	SystemCursorPointer
	SystemCursorNwResize
	SystemCursorNResize
	SystemCursorNeResize
	SystemCursorEResize
	SystemCursorSeResize
	SystemCursorSResize
	SystemCursorSwResize
	SystemCursorWResize
	SystemCursorCount
)

type MouseID uint32

type MouseButtonFlags uint32

const (
	ButtonLeft   MouseButtonFlags = 1
	ButtonMiddle MouseButtonFlags = 2
	ButtonRight  MouseButtonFlags = 3
	ButtonX1     MouseButtonFlags = 4
	ButtonX2     MouseButtonFlags = 5

	ButtonLmask  = 1 << (ButtonLeft - 1)
	ButtonMmask  = 1 << (ButtonMiddle - 1)
	ButtonRmask  = 1 << (ButtonRight - 1)
	ButtonX1mask = 1 << (ButtonX1 - 1)
	ButtonX2mask = 1 << (ButtonX2 - 1)
)

type Cursor struct{}

func CaptureMouse(enabled bool) bool {
	return sdlCaptureMouse(enabled)
}

func CreateColorCursor(surface *Surface, hotX int32, hotY int32) *Cursor {
	return sdlCreateColorCursor(surface, hotX, hotY)
}

func CreateCursor(data *uint8, mask *uint8, w int32, h int32, hotX int32, hotY int32) *Cursor {
	return sdlCreateCursor(data, mask, w, h, hotX, hotY)
}

func CreateSystemCursor(id SystemCursor) *Cursor {
	return sdlCreateSystemCursor(id)
}

// CursorVisible returns true if the cursor is being shown, or false if the cursor is hidden.
func CursorVisible() bool {
	return sdlCursorVisible()
}

func DestroyCursor(cursor *Cursor) {
	sdlDestroyCursor(cursor)
}

func GetCursor() *Cursor {
	return sdlGetCursor()
}

func GetDefaultCursor() *Cursor {
	return sdlGetDefaultCursor()
}

func GetGlobalMouseState(x *float32, y *float32) MouseButtonFlags {
	return sdlGetGlobalMouseState(x, y)
}

// GetMice returns a list of currently connected mice or nil on failure.
func GetMice() []MouseID {
	var count int32
	mice := sdlGetMice(&count)
	defer Free(unsafe.Pointer(mice))
	return mem.Copy(mice, count)
}

// GetMouseFocus returns the window with mouse focus.
func GetMouseFocus() *Window {
	return sdlGetMouseFocus()
}

// GetMouseNameForID returns the name of the selected mouse, or error on failure.
//
// This function returns "" if the mouse doesn't have a name.
func GetMouseNameForID(instanceId MouseID) (string, error) {
	ret := sdlGetMouseNameForID(instanceId)
	if ret == nil {
		return "", errors.New(GetError())
	}
	return convert.ToString(ret), nil
}

// GetMouseState queries SDL's cache for the synchronous mouse button state and the window-relative SDL-cursor position.
func GetMouseState(x *float32, y *float32) MouseButtonFlags {
	return sdlGetMouseState(x, y)
}

func GetRelativeMouseState(x *float32, y *float32) MouseButtonFlags {
	return sdlGetRelativeMouseState(x, y)
}

func GetWindowRelativeMouseMode(window *Window) bool {
	return sdlGetWindowRelativeMouseMode(window)
}

// HasMouse returns true if a mouse is connected, false otherwise.
func HasMouse() bool {
	return sdlHasMouse()
}

// HideCursor hides the cursor.
func HideCursor() bool {
	return sdlHideCursor()
}

func SetCursor(cursor *Cursor) bool {
	return sdlSetCursor(cursor)
}

func SetWindowRelativeMouseMode(window *Window, enabled bool) bool {
	return sdlSetWindowRelativeMouseMode(window, enabled)
}

// ShowCursor shows the cursor.
func ShowCursor() bool {
	return sdlShowCursor()
}

func WarpMouseGlobal(x float32, y float32) bool {
	return sdlWarpMouseGlobal(x, y)
}

// WarpMouseInWindow moves the mouse cursor to the given position within the window.
func WarpMouseInWindow(window *Window, x float32, y float32) {
	sdlWarpMouseInWindow(window, x, y)
}
