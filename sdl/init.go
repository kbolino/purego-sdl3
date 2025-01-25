package sdl

import (
	"runtime"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/shared"
)

var (
	sdlInit                    func(InitFlags) bool
	sdlQuit                    func()
	sdlGetError                func() string
	sdlCreateWindowAndRenderer func(string, int32, int32, WindowFlags, *Window, *Renderer) bool
	sdlSetRenderDrawColor      func(Renderer, uint8, uint8, uint8, uint8) bool
	sdlRenderClear             func(Renderer) bool
	sdlRenderPresent           func(Renderer) bool
	sdlDestroyRenderer         func(Renderer)
	sdlDestroyWindow           func(Window)
	sdlPollEvent               func(*Event) bool
	sdlSetHint                 func(string, string) bool
)

func init() {
	runtime.LockOSThread()

	var filename string
	switch runtime.GOOS {
	case "linux", "freebsd":
		filename = "libSDL3.so.0"
	case "windows":
		filename = "SDL3.dll"
	case "darwin":
		filename = "libSDL3.dylib"
	}

	lib := must(shared.Load(filename))

	purego.RegisterLibFunc(&sdlInit, lib, "SDL_Init")
	purego.RegisterLibFunc(&sdlQuit, lib, "SDL_Quit")
	purego.RegisterLibFunc(&sdlGetError, lib, "SDL_GetError")
	purego.RegisterLibFunc(&sdlCreateWindowAndRenderer, lib, "SDL_CreateWindowAndRenderer")
	purego.RegisterLibFunc(&sdlSetRenderDrawColor, lib, "SDL_SetRenderDrawColor")
	purego.RegisterLibFunc(&sdlRenderClear, lib, "SDL_RenderClear")
	purego.RegisterLibFunc(&sdlRenderPresent, lib, "SDL_RenderPresent")
	purego.RegisterLibFunc(&sdlDestroyRenderer, lib, "SDL_DestroyRenderer")
	purego.RegisterLibFunc(&sdlDestroyWindow, lib, "SDL_DestroyWindow")
	purego.RegisterLibFunc(&sdlPollEvent, lib, "SDL_PollEvent")
	purego.RegisterLibFunc(&sdlSetHint, lib, "SDL_SetHint")
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
