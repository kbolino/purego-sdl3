package sdl

import (
	"unsafe"
)

type Renderer unsafe.Pointer

// CreateWindowAndRenderer creates a window and default renderer.
func CreateWindowAndRenderer(title string, width, height int32, flags WindowFlags, window *Window, renderer *Renderer) bool {
	return sdlCreateWindowAndRenderer(title, width, height, flags, window, renderer)
}

// SetRenderDrawColor sets the color used for drawing operations.
func SetRenderDrawColor(renderer Renderer, r, g, b, a uint8) bool {
	return sdlSetRenderDrawColor(renderer, r, g, b, a)
}

// RenderPresent updates the screen with any rendering performed since the previous call.
func RenderPresent(renderer Renderer) bool {
	return sdlRenderPresent(renderer)
}

// RenderClear clears the current rendering target with the drawing color.
func RenderClear(renderer Renderer) bool {
	return sdlRenderClear(renderer)
}

// DestroyRenderer destroys the rendering context for a window and free all associated textures.
func DestroyRenderer(renderer Renderer) {
	sdlDestroyRenderer(renderer)
}
