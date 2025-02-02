package sdl

type HintPriority uint32

const (
	HintDefault HintPriority = iota
	HintNormal
	HintOverride
)

const (
	HintRenderDriver = "SDL_RENDER_DRIVER"
	HintRenderVsync  = "SDL_RENDER_VSYNC"
)

// SetHint sets a hint with normal priority.
func SetHint(name, value string) bool {
	return sdlSetHint(name, value)
}

// func AddHintCallback(name string, callback HintCallback, userdata unsafe.Pointer) bool {
//	return sdlAddHintCallback(name, callback, userdata)
// }

// func GetHint(name string) string {
//	return sdlGetHint(name)
// }

// func GetHintBoolean(name string, default_value bool) bool {
//	return sdlGetHintBoolean(name, default_value)
// }

// func RemoveHintCallback(name string, callback HintCallback, userdata unsafe.Pointer)  {
//	sdlRemoveHintCallback(name, callback, userdata)
// }

// func ResetHint(name string) bool {
//	return sdlResetHint(name)
// }

// func ResetHints()  {
//	sdlResetHints()
// }

// func SetHintWithPriority(name string, value string, priority HintPriority) bool {
//	return sdlSetHintWithPriority(name, value, priority)
// }
