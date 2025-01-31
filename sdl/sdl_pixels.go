package sdl

type PixelFormat uint32

const (
	PixelFormatUnknown      PixelFormat = 0
	PixelFormatIndex1lsb    PixelFormat = 0x11100100
	PixelFormatIndex1msb    PixelFormat = 0x11200100
	PixelFormatIndex2lsb    PixelFormat = 0x1C100200
	PixelFormatIndex2msb    PixelFormat = 0x1C200200
	PixelFormatIndex4lsb    PixelFormat = 0x12100400
	PixelFormatIndex4msb    PixelFormat = 0x12200400
	PixelFormatIndex8       PixelFormat = 0x13000801
	PixelFormatRgb332       PixelFormat = 0x14110801
	PixelFormatXrgb4444     PixelFormat = 0x15120C02
	PixelFormatXbgr4444     PixelFormat = 0x15520C02
	PixelFormatXrgb1555     PixelFormat = 0x15130F02
	PixelFormatXbgr1555     PixelFormat = 0x15530F02
	PixelFormatArgb4444     PixelFormat = 0x15321002
	PixelFormatRgba4444     PixelFormat = 0x15421002
	PixelFormatAbgr4444     PixelFormat = 0x15721002
	PixelFormatBgra4444     PixelFormat = 0x15821002
	PixelFormatArgb1555     PixelFormat = 0x15331002
	PixelFormatRgba5551     PixelFormat = 0x15441002
	PixelFormatAbgr1555     PixelFormat = 0x15731002
	PixelFormatBgra5551     PixelFormat = 0x15841002
	PixelFormatRgb565       PixelFormat = 0x15151002
	PixelFormatBgr565       PixelFormat = 0x15551002
	PixelFormatRgb24        PixelFormat = 0x17101803
	PixelFormatBgr24        PixelFormat = 0x17401803
	PixelFormatXrgb8888     PixelFormat = 0x16161804
	PixelFormatRgbx8888     PixelFormat = 0x16261804
	PixelFormatXbgr8888     PixelFormat = 0x16561804
	PixelFormatBgrx8888     PixelFormat = 0x16661804
	PixelFormatArgb8888     PixelFormat = 0x16362004
	PixelFormatRgba8888     PixelFormat = 0x16462004
	PixelFormatAbgr8888     PixelFormat = 0x16762004
	PixelFormatBgra8888     PixelFormat = 0x16862004
	PixelFormatXrgb2101010  PixelFormat = 0x16172004
	PixelFormatXbgr2101010  PixelFormat = 0x16572004
	PixelFormatArgb2101010  PixelFormat = 0x16372004
	PixelFormatAbgr2101010  PixelFormat = 0x16772004
	PixelFormatRgb48        PixelFormat = 0x18103006
	PixelFormatBgr48        PixelFormat = 0x18403006
	PixelFormatRgba64       PixelFormat = 0x18204008
	PixelFormatArgb64       PixelFormat = 0x18304008
	PixelFormatBgra64       PixelFormat = 0x18504008
	PixelFormatAbgr64       PixelFormat = 0x18604008
	PixelFormatRgb48Float   PixelFormat = 0x1A103006
	PixelFormatBgr48Float   PixelFormat = 0x1A403006
	PixelFormatRgba64Float  PixelFormat = 0x1A204008
	PixelFormatArgb64Float  PixelFormat = 0x1A304008
	PixelFormatBgra64Float  PixelFormat = 0x1A504008
	PixelFormatAbgr64Float  PixelFormat = 0x1A604008
	PixelFormatRgb96Float   PixelFormat = 0x1B10600C
	PixelFormatBgr96Float   PixelFormat = 0x1B40600C
	PixelFormatRgba128Float PixelFormat = 0x1B208010
	PixelFormatArgb128Float PixelFormat = 0x1B308010
	PixelFormatBgra128Float PixelFormat = 0x1B508010
	PixelFormatAbgr128Float PixelFormat = 0x1B608010
	PixelFormatYv12         PixelFormat = 0x32315659
	PixelFormatIyuv         PixelFormat = 0x56555949
	PixelFormatYuy2         PixelFormat = 0x32595559
	PixelFormatUyvy         PixelFormat = 0x59565955
	PixelFormatYvyu         PixelFormat = 0x55595659
	PixelFormatNv12         PixelFormat = 0x3231564E
	PixelFormatNv21         PixelFormat = 0x3132564E
	PixelFormatP010         PixelFormat = 0x30313050
	PixelFormatExternalOes  PixelFormat = 0x2053454F
)

// func CreatePalette(ncolors int32) *Palette {
//	return sdlCreatePalette(ncolors)
// }

// func DestroyPalette(palette *Palette)  {
//	sdlDestroyPalette(palette)
// }

// func GetMasksForPixelFormat(format PixelFormat, bpp *int32, Rmask *uint32, Gmask *uint32, Bmask *uint32, Amask *uint32) bool {
//	return sdlGetMasksForPixelFormat(format, bpp, Rmask, Gmask, Bmask, Amask)
// }

// func GetPixelFormatDetails(format PixelFormat) *PixelFormatDetails {
//	return sdlGetPixelFormatDetails(format)
// }

// func GetPixelFormatForMasks(bpp int32, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) PixelFormat {
//	return sdlGetPixelFormatForMasks(bpp, Rmask, Gmask, Bmask, Amask)
// }

// func GetPixelFormatName(format PixelFormat) string {
//	return sdlGetPixelFormatName(format)
// }

// func GetRGB(pixel uint32, format *PixelFormatDetails, palette *Palette, r *uint8, g *uint8, b *uint8)  {
//	sdlGetRGB(pixel, format, palette, r, g, b)
// }

// func GetRGBA(pixel uint32, format *PixelFormatDetails, palette *Palette, r *uint8, g *uint8, b *uint8, a *uint8)  {
//	sdlGetRGBA(pixel, format, palette, r, g, b, a)
// }

// func MapRGB(format *PixelFormatDetails, palette *Palette, r uint8, g uint8, b uint8) uint32 {
//	return sdlMapRGB(format, palette, r, g, b)
// }

// func MapRGBA(format *PixelFormatDetails, palette *Palette, r uint8, g uint8, b uint8, a uint8) uint32 {
//	return sdlMapRGBA(format, palette, r, g, b, a)
// }

// func SetPaletteColors(palette *Palette, colors *Color, firstcolor int32, ncolors int32) bool {
//	return sdlSetPaletteColors(palette, colors, firstcolor, ncolors)
// }
