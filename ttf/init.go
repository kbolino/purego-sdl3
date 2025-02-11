package ttf

import (
	"runtime"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/shared"
)

var (
	ttfVersion            func() int32
	ttfGetFreeTypeVersion func(*int32, *int32, *int32)
	ttfInit               func() bool
)

func init() {
	runtime.LockOSThread()

	var filename string
	switch runtime.GOOS {
	case "linux", "freebsd":
		filename = "libSDL3_ttf.so.0"
	case "windows":
		filename = "SDL3_ttf.dll"
	case "darwin":
		filename = "libSDL3_ttf.dylib"
	}

	lib, err := shared.Load(filename)
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&ttfVersion, lib, "TTF_Version")
	purego.RegisterLibFunc(&ttfGetFreeTypeVersion, lib, "TTF_GetFreeTypeVersion")
	purego.RegisterLibFunc(&ttfInit, lib, "TTF_Init")
}
