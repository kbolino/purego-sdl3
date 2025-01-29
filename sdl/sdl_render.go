package sdl

type Renderer struct{}

// Texture is a efficient driver-specific representation of pixel data.
type Texture struct {
	Format   PixelFormat
	W        int32
	H        int32
	Refcount int32
}

// CreateWindowAndRenderer creates a window and default renderer.
func CreateWindowAndRenderer(title string, width, height int32, flags WindowFlags, window **Window, renderer **Renderer) bool {
	return sdlCreateWindowAndRenderer(title, width, height, flags, window, renderer)
}

// SetRenderDrawColor sets the color used for drawing operations.
func SetRenderDrawColor(renderer *Renderer, r, g, b, a uint8) bool {
	return sdlSetRenderDrawColor(renderer, r, g, b, a)
}

// RenderPresent updates the screen with any rendering performed since the previous call.
func RenderPresent(renderer *Renderer) bool {
	return sdlRenderPresent(renderer)
}

// RenderClear clears the current rendering target with the drawing color.
func RenderClear(renderer *Renderer) bool {
	return sdlRenderClear(renderer)
}

// DestroyRenderer destroys the rendering context for a window and free all associated textures.
func DestroyRenderer(renderer *Renderer) {
	sdlDestroyRenderer(renderer)
}

// RenderRect draws a rectangle on the current rendering target at subpixel precision.
func RenderRect(renderer *Renderer, rect *FRect) bool {
	return sdlRenderRect(renderer, rect)
}

// RenderFillRect fills a rectangle on the current rendering target with the drawing color at subpixel precision.
func RenderFillRect(renderer *Renderer, rect *FRect) bool {
	return sdlRenderFillRect(renderer, rect)
}

// RenderDebugText draws debug text to a [Renderer].
func RenderDebugText(renderer *Renderer, x, y float32, str string) bool {
	return sdlRenderDebugText(renderer, x, y, str)
}

// CreateTextureFromSurface creates a texture from an existing surface.
func CreateTextureFromSurface(renderer *Renderer, surface *Surface) *Texture {
	return sdlCreateTextureFromSurface(renderer, surface)
}

// RenderTexture copies a portion of the texture to the current rendering target at subpixel precision.
func RenderTexture(renderer *Renderer, texture *Texture, srcrect *FRect, dstrect *FRect) bool {
	return sdlRenderTexture(renderer, texture, srcrect, dstrect)
}

// DestroyTexture destroys the specified texture.
func DestroyTexture(texture *Texture) {
	sdlDestroyTexture(texture)
}
