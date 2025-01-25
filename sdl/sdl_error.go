package sdl

// GetError retrieves a message about the last error that occurred on the current thread.
func GetError() string {
	return sdlGetError()
}
