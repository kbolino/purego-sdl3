package sdl

// GetRevision returns an arbitrary string, uniquely identifying the exact revision of the SDL library in use.
func GetRevision() string {
	return sdlGetRevision()
}

// GetVersion get the version of SDL that is linked against your program.
func GetVersion() (major, minor, patch int32) {
	version := sdlGetVersion()
	major = version / 1000000
	minor = version / 1000 % 1000
	patch = version % 1000
	return
}
