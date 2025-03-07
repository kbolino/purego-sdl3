package sdl

type DateFormat uint32

const (
	DateFormatYYYYMMDD DateFormat = iota
	DateFormatDDMMYYYY
	DateFormatMMDDYYYY
)

type TimeFormat uint32

const (
	TimeFormat24HR TimeFormat = iota
	TimeFormat12HR
)

// func DateTimeToTime(dt *DateTime, ticks *Time) bool {
//	return sdlDateTimeToTime(dt, ticks)
// }

// func GetCurrentTime(ticks *Time) bool {
//	return sdlGetCurrentTime(ticks)
// }

// func GetDateTimeLocalePreferences(dateFormat *DateFormat, timeFormat *TimeFormat) bool {
//	return sdlGetDateTimeLocalePreferences(dateFormat, timeFormat)
// }

// func GetDayOfWeek(year int32, month int32, day int32) int32 {
//	return sdlGetDayOfWeek(year, month, day)
// }

// func GetDayOfYear(year int32, month int32, day int32) int32 {
//	return sdlGetDayOfYear(year, month, day)
// }

// func GetDaysInMonth(year int32, month int32) int32 {
//	return sdlGetDaysInMonth(year, month)
// }

// func TimeFromWindows(dwLowDateTime uint32, dwHighDateTime uint32) Time {
//	return sdlTimeFromWindows(dwLowDateTime, dwHighDateTime)
// }

// func TimeToDateTime(ticks Time, dt *DateTime, localTime bool) bool {
//	return sdlTimeToDateTime(ticks, dt, localTime)
// }

// func TimeToWindows(ticks Time, dwLowDateTime *uint32, dwHighDateTime *uint32)  {
//	sdlTimeToWindows(ticks, dwLowDateTime, dwHighDateTime)
// }
