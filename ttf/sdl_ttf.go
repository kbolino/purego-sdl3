package ttf

// Version gets the version of the dynamically linked SDL_ttf library.
func Version() (major, minor, patch int32) {
	version := ttfVersion()
	major = version / 1000000
	minor = version / 1000 % 1000
	patch = version % 1000
	return
}

// GetFreeTypeVersion queries the version of the FreeType library in use.
func GetFreeTypeVersion() (major, minor, patch int32) {
	ttfGetFreeTypeVersion(&major, &minor, &patch)
	return
}

// Init initializes SDL_ttf.
func Init() bool {
	return ttfInit()
}

// GetHarfBuzzVersion queries the version of the HarfBuzz library in use.
func GetHarfBuzzVersion() (major, minor, patch int32) {
	ttfGetHarfBuzzVersion(&major, &minor, &patch)
	return
}
