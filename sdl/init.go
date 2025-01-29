package sdl

import (
	"runtime"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/shared"
)

var (
	sdlInit                     func(InitFlags) bool
	sdlQuitSubSystem            func(InitFlags)
	sdlQuit                     func()
	sdlGetError                 func() string
	sdlClearError               func() bool
	sdlCreateWindowAndRenderer  func(string, int32, int32, WindowFlags, **Window, **Renderer) bool
	sdlSetRenderDrawColor       func(*Renderer, uint8, uint8, uint8, uint8) bool
	sdlRenderClear              func(*Renderer) bool
	sdlRenderPresent            func(*Renderer) bool
	sdlDestroyRenderer          func(*Renderer)
	sdlDestroyWindow            func(*Window)
	sdlPollEvent                func(*Event) bool
	sdlSetHint                  func(string, string) bool
	sdlGetPowerInfo             func(*int32, *int32) PowerState
	sdlRenderRect               func(*Renderer, *FRect) bool
	sdlRenderFillRect           func(*Renderer, *FRect) bool
	sdlRenderDebugText          func(*Renderer, float32, float32, string) bool
	sdlGetPreferredLocales      func(*int32) **Locale
	sdlIOFromConstMem           func([]byte, int) *IOStream
	sdlCloseIO                  func(*IOStream) bool
	sdlLoadBMPIO                func(*IOStream, bool) *Surface
	sdlDestroySurface           func(*Surface)
	sdlCreateTextureFromSurface func(*Renderer, *Surface) *Texture
	sdlRenderTexture            func(*Renderer, *Texture, *FRect, *FRect) bool
	sdlDestroyTexture           func(*Texture)
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

	lib, err := shared.Load(filename)
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&sdlInit, lib, "SDL_Init")
	purego.RegisterLibFunc(&sdlQuitSubSystem, lib, "SDL_QuitSubSystem")
	purego.RegisterLibFunc(&sdlQuit, lib, "SDL_Quit")
	purego.RegisterLibFunc(&sdlGetError, lib, "SDL_GetError")
	purego.RegisterLibFunc(&sdlClearError, lib, "SDL_ClearError")
	purego.RegisterLibFunc(&sdlCreateWindowAndRenderer, lib, "SDL_CreateWindowAndRenderer")
	purego.RegisterLibFunc(&sdlSetRenderDrawColor, lib, "SDL_SetRenderDrawColor")
	purego.RegisterLibFunc(&sdlRenderClear, lib, "SDL_RenderClear")
	purego.RegisterLibFunc(&sdlRenderPresent, lib, "SDL_RenderPresent")
	purego.RegisterLibFunc(&sdlDestroyRenderer, lib, "SDL_DestroyRenderer")
	purego.RegisterLibFunc(&sdlDestroyWindow, lib, "SDL_DestroyWindow")
	purego.RegisterLibFunc(&sdlPollEvent, lib, "SDL_PollEvent")
	purego.RegisterLibFunc(&sdlSetHint, lib, "SDL_SetHint")
	purego.RegisterLibFunc(&sdlGetPowerInfo, lib, "SDL_GetPowerInfo")
	purego.RegisterLibFunc(&sdlRenderRect, lib, "SDL_RenderRect")
	purego.RegisterLibFunc(&sdlRenderFillRect, lib, "SDL_RenderFillRect")
	purego.RegisterLibFunc(&sdlRenderDebugText, lib, "SDL_RenderDebugText")
	purego.RegisterLibFunc(&sdlGetPreferredLocales, lib, "SDL_GetPreferredLocales")
	purego.RegisterLibFunc(&sdlIOFromConstMem, lib, "SDL_IOFromConstMem")
	purego.RegisterLibFunc(&sdlCloseIO, lib, "SDL_CloseIO")
	purego.RegisterLibFunc(&sdlLoadBMPIO, lib, "SDL_LoadBMP_IO")
	purego.RegisterLibFunc(&sdlDestroySurface, lib, "SDL_DestroySurface")
	purego.RegisterLibFunc(&sdlCreateTextureFromSurface, lib, "SDL_CreateTextureFromSurface")
	purego.RegisterLibFunc(&sdlRenderTexture, lib, "SDL_RenderTexture")
	purego.RegisterLibFunc(&sdlDestroyTexture, lib, "SDL_DestroyTexture")
}
