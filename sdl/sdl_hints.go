package sdl

const (
	HintRenderDriver = "SDL_RENDER_DRIVER"
	HintRenderVsync  = "SDL_RENDER_VSYNC"
)

// SetHint sets a hint with normal priority.
func SetHint(name, value string) bool {
	return sdlSetHint(name, value)
}
