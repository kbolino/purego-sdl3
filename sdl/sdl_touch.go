package sdl

type TouchDeviceType int32

const (
	TouchDeviceInvalid TouchDeviceType = iota - 1
	TouchDeviceDirect
	TouchDeviceIndirectAbsolute
	TouchDeviceIndirectRelative
)

type TouchID uint64

type FingerID uint64

// func GetTouchDeviceName(touchID TouchID) string {
//	return sdlGetTouchDeviceName(touchID)
// }

// func GetTouchDevices(count *int32) *TouchID {
//	return sdlGetTouchDevices(count)
// }

// func GetTouchDeviceType(touchID TouchID) TouchDeviceType {
//	return sdlGetTouchDeviceType(touchID)
// }

// func GetTouchFingers(touchID TouchID, count *int32) **Finger {
//	return sdlGetTouchFingers(touchID, count)
// }
