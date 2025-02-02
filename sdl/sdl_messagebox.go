package sdl

type MessageBoxColorType uint32

const (
	MessageBoxColorBackground MessageBoxColorType = iota
	MessageBoxColorText
	MessageBoxColorButtonBorder
	MessageBoxColorButtonBackground
	MessageBoxColorButtonSelected
	MessageBoxColorCount
)

// func ShowMessageBox(messageboxdata *MessageBoxData, buttonid *int32) bool {
//	return sdlShowMessageBox(messageboxdata, buttonid)
// }

// func ShowSimpleMessageBox(flags MessageBoxFlags, title string, message string, window *Window) bool {
//	return sdlShowSimpleMessageBox(flags, title, message, window)
// }
