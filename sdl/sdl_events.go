package sdl

import (
	"unsafe"
)

type EventType uint32

const (
	EventFirst   EventType = 0
	EventQuit    EventType = 0x100
	EventKeyDown EventType = 0x300
)

type Event [128]byte

func (e *Event) Type() EventType {
	return *(*EventType)(unsafe.Pointer(e))
}

func (e *Event) Key() KeyboardEvent {
	return *(*KeyboardEvent)(unsafe.Pointer(e))
}

type KeyboardEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Which     KeyboardID
	Scancode  Scancode
	Key       Keycode
	Mod       Keymod
	Raw       uint16
	Down      bool
	Repeat    bool
}

// PollEvent polls for currently pending events.
func PollEvent(event *Event) bool {
	return sdlPollEvent(event)
}
