package sdl

import (
	"unsafe"
)

type EventType uint32

const (
	EventFirst           EventType = 0
	EventQuit            EventType = 0x100
	EventKeyDown         EventType = 0x300
	EventKeyUp           EventType = 0x301
	EventMouseMotion     EventType = 0x400
	EventMouseButtonDown EventType = 0x401
	EventMouseButtonUp   EventType = 0x402
)

type Event [128]byte

func (e *Event) Type() EventType {
	return *(*EventType)(unsafe.Pointer(e))
}

func (e *Event) Key() KeyboardEvent {
	return *(*KeyboardEvent)(unsafe.Pointer(e))
}

func (e *Event) MouseMotion() MouseMotionEvent {
	return *(*MouseMotionEvent)(unsafe.Pointer(e))
}

func (e *Event) MouseButton() MouseButtonEvent {
	return *(*MouseButtonEvent)(unsafe.Pointer(e))
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

type MouseMotionEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Which     MouseID
	State     MouseButtonFlags
	X         float32
	Y         float32
	Xrel      float32
	Yrel      float32
}

type MouseButtonEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Which     MouseID
	Button    uint8
	Down      bool
	Clicks    uint8
	Padding   uint8
	X         float32
	Y         float32
}

// PollEvent polls for currently pending events.
func PollEvent(event *Event) bool {
	return sdlPollEvent(event)
}

// func AddEventWatch(filter EventFilter, userdata unsafe.Pointer) bool {
//	return sdlAddEventWatch(filter, userdata)
// }

// func EventEnabled(type uint32) bool {
//	return sdlEventEnabled(type)
// }

// func FilterEvents(filter EventFilter, userdata unsafe.Pointer)  {
//	sdlFilterEvents(filter, userdata)
// }

// func FlushEvent(type uint32)  {
//	sdlFlushEvent(type)
// }

// func FlushEvents(minType uint32, maxType uint32)  {
//	sdlFlushEvents(minType, maxType)
// }

// func GetEventFilter(filter *EventFilter, userdata *unsafe.Pointer) bool {
//	return sdlGetEventFilter(filter, userdata)
// }

// func GetWindowFromEvent(event *Event) *Window {
//	return sdlGetWindowFromEvent(event)
// }

// func HasEvent(type uint32) bool {
//	return sdlHasEvent(type)
// }

// func HasEvents(minType uint32, maxType uint32) bool {
//	return sdlHasEvents(minType, maxType)
// }

// func PeepEvents(events *Event, numevents int32, action EventAction, minType uint32, maxType uint32) int32 {
//	return sdlPeepEvents(events, numevents, action, minType, maxType)
// }

// PumpEvents updates the event queue and internal input device state.
func PumpEvents() {
	sdlPumpEvents()
}

// func PushEvent(event *Event) bool {
//	return sdlPushEvent(event)
// }

// func RegisterEvents(numevents int32) uint32 {
//	return sdlRegisterEvents(numevents)
// }

// func RemoveEventWatch(filter EventFilter, userdata unsafe.Pointer)  {
//	sdlRemoveEventWatch(filter, userdata)
// }

// func SetEventEnabled(type uint32, enabled bool)  {
//	sdlSetEventEnabled(type, enabled)
// }

// func SetEventFilter(filter EventFilter, userdata unsafe.Pointer)  {
//	sdlSetEventFilter(filter, userdata)
// }

// func WaitEvent(event *Event) bool {
//	return sdlWaitEvent(event)
// }

// func WaitEventTimeout(event *Event, timeoutMS int32) bool {
//	return sdlWaitEventTimeout(event, timeoutMS)
// }
