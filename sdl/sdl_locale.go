package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/convert"
	"github.com/jupiterrider/purego-sdl3/internal/mem"
)

type Locale struct {
	language, country *byte
}

func (l *Locale) Language() string {
	return convert.ToString(l.language)
}

func (l *Locale) Country() string {
	return convert.ToString(l.country)
}

// GetPreferredLocales reports the user's preferred locale.
func GetPreferredLocales() []*Locale {
	var count int32
	locales := sdlGetPreferredLocales(&count)
	if locales == nil {
		return nil
	}
	defer Free(unsafe.Pointer(locales))
	return mem.Copy(locales, count)
}
