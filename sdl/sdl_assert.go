package sdl

type AssertState uint32

const (
	AssertionRetry AssertState = iota
	AssertionBreak
	AssertionAbort
	AssertionIgnore
	AssertionAlwaysIgnore
)

// func GetAssertionHandler(puserdata *unsafe.Pointer) AssertionHandler {
//	return sdlGetAssertionHandler(puserdata)
// }

// func GetAssertionReport() *AssertData {
//	return sdlGetAssertionReport()
// }

// func GetDefaultAssertionHandler() AssertionHandler {
//	return sdlGetDefaultAssertionHandler()
// }

// func ReportAssertion(data *AssertData, func string, file string, line int32) AssertState {
//	return sdlReportAssertion(data, func, file, line)
// }

// func ResetAssertionReport()  {
//	sdlResetAssertionReport()
// }

// func SetAssertionHandler(handler AssertionHandler, userdata unsafe.Pointer)  {
//	sdlSetAssertionHandler(handler, userdata)
// }
