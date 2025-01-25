package sdl

type InitFlags uint32

const (
	InitAudio    InitFlags = 0x00000010
	InitVideo    InitFlags = 0x00000020
	InitJoystick InitFlags = 0x00000200
	InitHaptic   InitFlags = 0x00001000
	InitGamepad  InitFlags = 0x00002000
	InitEvents   InitFlags = 0x00004000
	InitSensor   InitFlags = 0x00008000
	InitCamera   InitFlags = 0x00010000
)

// Init initializes the SDL library.
func Init(flags InitFlags) bool {
	return sdlInit(flags)
}

// Quit cleans up all initialized subsystems.
func Quit() {
	sdlQuit()
}
