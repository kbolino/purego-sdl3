package sdl

// PowerState is the basic state for the system's power supply return by [GetPowerInfo].
type PowerState int32

const (
	PowerStateError PowerState = iota - 1
	PowerStateUnknown
	PowerStateOnBattery
	PowerStateNoBattery
	PowerStateCharging
	PowerStateCharged
)

// GetPowerInfo gets the current power supply details.
func GetPowerInfo(seconds *int32, percent *int32) PowerState {
	return sdlGetPowerInfo(seconds, percent)
}
