package sdl

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/convert"
)

type EventAction uint32

const (
	AddEvent EventAction = iota
	PeekEvent
	GetEvent
)

type EventType uint32

const (
	EventFirst                      EventType = 0x0
	EventQuit                       EventType = 0x100
	EventTerminating                EventType = 0x101
	EventLowMemory                  EventType = 0x102
	EventWillEnterBackground        EventType = 0x103
	EventDidEnterBackground         EventType = 0x104
	EventWillEnterForeground        EventType = 0x105
	EventDidEnterForeground         EventType = 0x106
	EventLocaleChanged              EventType = 0x107
	EventSystemThemeChanged         EventType = 0x108
	EventDisplayOrientation         EventType = 0x151
	EventDisplayFirst               EventType = 0x151
	EventDisplayAdded               EventType = 0x152
	EventDisplayRemoved             EventType = 0x153
	EventDisplayMoved               EventType = 0x154
	EventDisplayDesktopModeChanged  EventType = 0x155
	EventDisplayCurrentModeChanged  EventType = 0x156
	EventDisplayContentScaleChanged EventType = 0x157
	EventDisplayLast                EventType = 0x157
	EventWindowShown                EventType = 0x202
	EventWindowFirst                EventType = 0x202
	EventWindowHidden               EventType = 0x203
	EventWindowExposed              EventType = 0x204
	EventWindowMoved                EventType = 0x205
	EventWindowResized              EventType = 0x206
	EventWindowPixelSizeChanged     EventType = 0x207
	EventWindowMetalViewResized     EventType = 0x208
	EventWindowMinimized            EventType = 0x209
	EventWindowMaximized            EventType = 0x20A
	EventWindowRestored             EventType = 0x20B
	EventWindowMouseEnter           EventType = 0x20C
	EventWindowMouseLeave           EventType = 0x20D
	EventWindowFocusGained          EventType = 0x20E
	EventWindowFocusLost            EventType = 0x20F
	EventWindowCloseRequested       EventType = 0x210
	EventWindowHitTest              EventType = 0x211
	EventWindowIccprofChanged       EventType = 0x212
	EventWindowDisplayChanged       EventType = 0x213
	EventWindowDisplayScaleChanged  EventType = 0x214
	EventWindowSafeAreaChanged      EventType = 0x215
	EventWindowOccluded             EventType = 0x216
	EventWindowEnterFullscreen      EventType = 0x217
	EventWindowLeaveFullscreen      EventType = 0x218
	EventWindowDestroyed            EventType = 0x219
	EventWindowHdrStateChanged      EventType = 0x21A
	EventWindowLast                 EventType = 0x21A
	EventKeyDown                    EventType = 0x300
	EventKeyUp                      EventType = 0x301
	EventTextEditing                EventType = 0x302
	EventTextInput                  EventType = 0x303
	EventKeymapChanged              EventType = 0x304
	EventKeyboardAdded              EventType = 0x305
	EventKeyboardRemoved            EventType = 0x306
	EventTextEditingCandidates      EventType = 0x307
	EventMouseMotion                EventType = 0x400
	EventMouseButtonDown            EventType = 0x401
	EventMouseButtonUp              EventType = 0x402
	EventMouseWheel                 EventType = 0x403
	EventMouseAdded                 EventType = 0x404
	EventMouseRemoved               EventType = 0x405
	EventJoystickAxisMotion         EventType = 0x600
	EventJoystickBallMotion         EventType = 0x601
	EventJoystickHatMotion          EventType = 0x602
	EventJoystickButtonDown         EventType = 0x603
	EventJoystickButtonUp           EventType = 0x604
	EventJoystickAdded              EventType = 0x605
	EventJoystickRemoved            EventType = 0x606
	EventJoystickBatteryUpdated     EventType = 0x607
	EventJoystickUpdateComplete     EventType = 0x608
	EventGamepadAxisMotion          EventType = 0x650
	EventGamepadButtonDown          EventType = 0x651
	EventGamepadButtonUp            EventType = 0x652
	EventGamepadAdded               EventType = 0x653
	EventGamepadRemoved             EventType = 0x654
	EventGamepadRemapped            EventType = 0x655
	EventGamepadTouchpadDown        EventType = 0x656
	EventGamepadTouchpadMotion      EventType = 0x657
	EventGamepadTouchpadUp          EventType = 0x658
	EventGamepadSensorUpdate        EventType = 0x659
	EventGamepadUpdateComplete      EventType = 0x65A
	EventGamepadSteamHandleUpdated  EventType = 0x65B
	EventFingerDown                 EventType = 0x700
	EventFingerUp                   EventType = 0x701
	EventFingerMotion               EventType = 0x702
	EventFingerCanceled             EventType = 0x703
	EventClipboardUpdate            EventType = 0x900
	EventDropFile                   EventType = 0x1000
	EventDropText                   EventType = 0x1001
	EventDropBegin                  EventType = 0x1002
	EventDropComplete               EventType = 0x1003
	EventDropPosition               EventType = 0x1004
	EventAudioDeviceAdded           EventType = 0x1100
	EventAudioDeviceRemoved         EventType = 0x1101
	EventAudioDeviceFormatChanged   EventType = 0x1102
	EventSensorUpdate               EventType = 0x1200
	EventPenProximityIn             EventType = 0x1300
	EventPenProximityOut            EventType = 0x1301
	EventPenDown                    EventType = 0x1302
	EventPenUp                      EventType = 0x1303
	EventPenButtonDown              EventType = 0x1304
	EventPenButtonUp                EventType = 0x1305
	EventPenMotion                  EventType = 0x1306
	EventPenAxis                    EventType = 0x1307
	EventCameraDeviceAdded          EventType = 0x1400
	EventCameraDeviceRemoved        EventType = 0x1401
	EventCameraDeviceApproved       EventType = 0x1402
	EventCameraDeviceDenied         EventType = 0x1403
	EventRenderTargetsReset         EventType = 0x2000
	EventRenderDeviceReset          EventType = 0x2001
	EventRenderDeviceLost           EventType = 0x2002
	EventPrivate0                   EventType = 0x4000
	EventPrivate1                   EventType = 0x4001
	EventPrivate2                   EventType = 0x4002
	EventPrivate3                   EventType = 0x4003
	EventPollSentinel               EventType = 0x7F00
	EventUser                       EventType = 0x8000
	EventLast                       EventType = 0xFFFF
	EventEnumPadding                EventType = 0x7FFFFFFF
)

type Event [128]byte

func (e *Event) Type() EventType {
	return *(*EventType)(unsafe.Pointer(e))
}

func (e *Event) Key() KeyboardEvent {
	return *(*KeyboardEvent)(unsafe.Pointer(e))
}

func (e *Event) Text() TextInputEvent {
	return *(*TextInputEvent)(unsafe.Pointer(e))
}

func (e *Event) Motion() MouseMotionEvent {
	return *(*MouseMotionEvent)(unsafe.Pointer(e))
}

func (e *Event) Button() MouseButtonEvent {
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

type TextInputEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	text      *byte
}

func (t *TextInputEvent) Text() string {
	return convert.ToString(t.text)
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

type EventFilter func(userdata unsafe.Pointer, event *Event) bool

func (e EventFilter) toCallback() uintptr {
	// workaround to avoid panic "expected function with one uintptr-sized result" on Windows
	return purego.NewCallback(func(userdata unsafe.Pointer, event *Event) uintptr {
		if e(userdata, event) {
			return 0
		} else {
			return 1
		}
	})
}

// PollEvent polls for currently pending events.
func PollEvent(event *Event) bool {
	return sdlPollEvent(event)
}

// AddEventWatch adds a callback to be triggered when an event is added to the event queue.
func AddEventWatch(filter EventFilter, userdata unsafe.Pointer) bool {
	return sdlAddEventWatch(filter.toCallback(), userdata)
}

// EventEnabled returns true if the event is being processed, false otherwise.
func EventEnabled(eventType EventType) bool {
	return sdlEventEnabled(eventType)
}

func FilterEvents(filter EventFilter, userdata unsafe.Pointer) {
	sdlFilterEvents(filter.toCallback(), userdata)
}

// FlushEvent clears events of a specific type from the event queue.
func FlushEvent(eventType EventType) {
	sdlFlushEvent(eventType)
}

// FlushEvents clears events of a range of types from the event queue.
func FlushEvents(minType, maxType EventType) {
	sdlFlushEvents(minType, maxType)
}

// func GetEventFilter(filter *EventFilter, userdata *unsafe.Pointer) bool {
//	return sdlGetEventFilter(filter, userdata)
// }

// GetWindowFromEvent returns the associated window with an event or nil if there is none.
func GetWindowFromEvent(event *Event) *Window {
	return sdlGetWindowFromEvent(event)
}

// HasEvent checks for the existence of a certain event type in the event queue.
func HasEvent(eventType EventType) bool {
	return sdlHasEvent(eventType)
}

// HasEvents checks for the existence of certain event types in the event queue.
//
// Returns true if events with type >= minType and <= maxType are present, or false if not.
func HasEvents(minType, maxType EventType) bool {
	return sdlHasEvents(minType, maxType)
}

// func PeepEvents(events *Event, numevents int32, action EventAction, minType uint32, maxType uint32) int32 {
//	return sdlPeepEvents(events, numevents, action, minType, maxType)
// }

// PumpEvents updates the event queue and internal input device state.
func PumpEvents() {
	sdlPumpEvents()
}

// PushEvent adds an event to the event queue.
func PushEvent(event *Event) bool {
	return sdlPushEvent(event)
}

// func RegisterEvents(numevents int32) uint32 {
//	return sdlRegisterEvents(numevents)
// }

// func RemoveEventWatch(filter EventFilter, userdata unsafe.Pointer)  {
//	sdlRemoveEventWatch(filter, userdata)
// }

// SetEventEnabled sets the state of processing events by type.
func SetEventEnabled(eventType EventType, enabled bool) {
	sdlSetEventEnabled(eventType, enabled)
}

func SetEventFilter(filter EventFilter, userdata unsafe.Pointer) {
	sdlSetEventFilter(filter.toCallback(), userdata)
}

// WaitEvent waits indefinitely for the next available event.
func WaitEvent(event *Event) bool {
	return sdlWaitEvent(event)
}

// WaitEventTimeout waits until the specified timeout (in milliseconds) for the next available event.
func WaitEventTimeout(event *Event, timeoutMS int32) bool {
	return sdlWaitEventTimeout(event, timeoutMS)
}
