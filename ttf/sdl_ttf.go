package ttf

// Version gets the version of the dynamically linked SDL_ttf library.
func Version() (major, minor, micro int32) {
	version := ttfVersion()
	major = version / 1000000
	minor = version / 1000 % 1000
	micro = version % 1000
	return
}

// GetFreeTypeVersion queries the version of the FreeType library in use.
func GetFreeTypeVersion() (major, minor, micro int32) {
	ttfGetFreeTypeVersion(&major, &minor, &micro)
	return
}

// Init initializes SDL_ttf.
func Init() bool {
	return ttfInit()
}
