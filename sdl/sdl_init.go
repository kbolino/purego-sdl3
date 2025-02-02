package sdl

type AppResult uint32

const (
	AppContinue AppResult = iota
	AppSuccess
	AppFailure
)

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

// InitSubSystem is a compatibility function to initialize the SDL library.
//
// This function and [Init] are interchangeable.
func InitSubSystem(flags InitFlags) bool {
	return sdlInit(flags)
}

// Quit cleans up all initialized subsystems.
func Quit() {
	sdlQuit()
}

// QuitSubSystem shuts down specific SDL subsystems.
//
// You still need to call [Quit] even if you close all open subsystems with this function.
func QuitSubSystem(flags InitFlags) {
	sdlQuitSubSystem(flags)
}

// GetAppMetadataProperty gets metadata about your app.
func GetAppMetadataProperty(name string) string {
	return sdlGetAppMetadataProperty(name)
}

// IsMainThread returns true if this is the main thread.
func IsMainThread() bool {
	return sdlIsMainThread()
}

// func RunOnMainThread(callback MainThreadCallback, userdata unsafe.Pointer, wait_complete bool) bool {
//	return sdlRunOnMainThread(callback, userdata, wait_complete)
// }

// func SetAppMetadata(appname string, appversion string, appidentifier string) bool {
//	return sdlSetAppMetadata(appname, appversion, appidentifier)
// }

// func SetAppMetadataProperty(name string, value string) bool {
//	return sdlSetAppMetadataProperty(name, value)
// }

// func WasInit(flags InitFlags) InitFlags {
//	return sdlWasInit(flags)
// }
