package sdl

// func AttachVirtualJoystick(desc *VirtualJoystickDesc) JoystickID {
//	return sdlAttachVirtualJoystick(desc)
// }

// func CloseJoystick(joystick *Joystick)  {
//	sdlCloseJoystick(joystick)
// }

// func DetachVirtualJoystick(instance_id JoystickID) bool {
//	return sdlDetachVirtualJoystick(instance_id)
// }

// func GetJoystickAxis(joystick *Joystick, axis int32) int16 {
//	return sdlGetJoystickAxis(joystick, axis)
// }

// func GetJoystickAxisInitialState(joystick *Joystick, axis int32, state *int16) bool {
//	return sdlGetJoystickAxisInitialState(joystick, axis, state)
// }

// func GetJoystickBall(joystick *Joystick, ball int32, dx *int32, dy *int32) bool {
//	return sdlGetJoystickBall(joystick, ball, dx, dy)
// }

// func GetJoystickButton(joystick *Joystick, button int32) bool {
//	return sdlGetJoystickButton(joystick, button)
// }

// func GetJoystickConnectionState(joystick *Joystick) JoystickConnectionState {
//	return sdlGetJoystickConnectionState(joystick)
// }

// func GetJoystickFirmwareVersion(joystick *Joystick) uint16 {
//	return sdlGetJoystickFirmwareVersion(joystick)
// }

// func GetJoystickFromID(instance_id JoystickID) *Joystick {
//	return sdlGetJoystickFromID(instance_id)
// }

// func GetJoystickFromPlayerIndex(player_index int32) *Joystick {
//	return sdlGetJoystickFromPlayerIndex(player_index)
// }

// func GetJoystickGUID(joystick *Joystick) GUID {
//	return sdlGetJoystickGUID(joystick)
// }

// func GetJoystickGUIDForID(instance_id JoystickID) GUID {
//	return sdlGetJoystickGUIDForID(instance_id)
// }

// func GetJoystickGUIDInfo(guid GUID, vendor *uint16, product *uint16, version *uint16, crc16 *uint16)  {
//	sdlGetJoystickGUIDInfo(guid, vendor, product, version, crc16)
// }

// func GetJoystickHat(joystick *Joystick, hat int32) uint8 {
//	return sdlGetJoystickHat(joystick, hat)
// }

// func GetJoystickID(joystick *Joystick) JoystickID {
//	return sdlGetJoystickID(joystick)
// }

// func GetJoystickName(joystick *Joystick) string {
//	return sdlGetJoystickName(joystick)
// }

// func GetJoystickNameForID(instance_id JoystickID) string {
//	return sdlGetJoystickNameForID(instance_id)
// }

// func GetJoystickPath(joystick *Joystick) string {
//	return sdlGetJoystickPath(joystick)
// }

// func GetJoystickPathForID(instance_id JoystickID) string {
//	return sdlGetJoystickPathForID(instance_id)
// }

// func GetJoystickPlayerIndex(joystick *Joystick) int32 {
//	return sdlGetJoystickPlayerIndex(joystick)
// }

// func GetJoystickPlayerIndexForID(instance_id JoystickID) int32 {
//	return sdlGetJoystickPlayerIndexForID(instance_id)
// }

// func GetJoystickPowerInfo(joystick *Joystick, percent *int32) PowerState {
//	return sdlGetJoystickPowerInfo(joystick, percent)
// }

// func GetJoystickProduct(joystick *Joystick) uint16 {
//	return sdlGetJoystickProduct(joystick)
// }

// func GetJoystickProductForID(instance_id JoystickID) uint16 {
//	return sdlGetJoystickProductForID(instance_id)
// }

// func GetJoystickProductVersion(joystick *Joystick) uint16 {
//	return sdlGetJoystickProductVersion(joystick)
// }

// func GetJoystickProductVersionForID(instance_id JoystickID) uint16 {
//	return sdlGetJoystickProductVersionForID(instance_id)
// }

// func GetJoystickProperties(joystick *Joystick) PropertiesID {
//	return sdlGetJoystickProperties(joystick)
// }

// func GetJoysticks(count *int32) *JoystickID {
//	return sdlGetJoysticks(count)
// }

// func GetJoystickSerial(joystick *Joystick) string {
//	return sdlGetJoystickSerial(joystick)
// }

// func GetJoystickType(joystick *Joystick) JoystickType {
//	return sdlGetJoystickType(joystick)
// }

// func GetJoystickTypeForID(instance_id JoystickID) JoystickType {
//	return sdlGetJoystickTypeForID(instance_id)
// }

// func GetJoystickVendor(joystick *Joystick) uint16 {
//	return sdlGetJoystickVendor(joystick)
// }

// func GetJoystickVendorForID(instance_id JoystickID) uint16 {
//	return sdlGetJoystickVendorForID(instance_id)
// }

// func GetNumJoystickAxes(joystick *Joystick) int32 {
//	return sdlGetNumJoystickAxes(joystick)
// }

// func GetNumJoystickBalls(joystick *Joystick) int32 {
//	return sdlGetNumJoystickBalls(joystick)
// }

// func GetNumJoystickButtons(joystick *Joystick) int32 {
//	return sdlGetNumJoystickButtons(joystick)
// }

// func GetNumJoystickHats(joystick *Joystick) int32 {
//	return sdlGetNumJoystickHats(joystick)
// }

// func HasJoystick() bool {
//	return sdlHasJoystick()
// }

// func IsJoystickVirtual(instance_id JoystickID) bool {
//	return sdlIsJoystickVirtual(instance_id)
// }

// func JoystickConnected(joystick *Joystick) bool {
//	return sdlJoystickConnected(joystick)
// }

// func JoystickEventsEnabled() bool {
//	return sdlJoystickEventsEnabled()
// }

// func LockJoysticks()  {
//	sdlLockJoysticks()
// }

// func OpenJoystick(instance_id JoystickID) *Joystick {
//	return sdlOpenJoystick(instance_id)
// }

// func RumbleJoystick(joystick *Joystick, low_frequency_rumble uint16, high_frequency_rumble uint16, duration_ms uint32) bool {
//	return sdlRumbleJoystick(joystick, low_frequency_rumble, high_frequency_rumble, duration_ms)
// }

// func RumbleJoystickTriggers(joystick *Joystick, left_rumble uint16, right_rumble uint16, duration_ms uint32) bool {
//	return sdlRumbleJoystickTriggers(joystick, left_rumble, right_rumble, duration_ms)
// }

// func SendJoystickEffect(joystick *Joystick, data unsafe.Pointer, size int32) bool {
//	return sdlSendJoystickEffect(joystick, data, size)
// }

// func SendJoystickVirtualSensorData(joystick *Joystick, type SensorType, sensor_timestamp uint64, data *float32, num_values int32) bool {
//	return sdlSendJoystickVirtualSensorData(joystick, type, sensor_timestamp, data, num_values)
// }

// func SetJoystickEventsEnabled(enabled bool)  {
//	sdlSetJoystickEventsEnabled(enabled)
// }

// func SetJoystickLED(joystick *Joystick, red uint8, green uint8, blue uint8) bool {
//	return sdlSetJoystickLED(joystick, red, green, blue)
// }

// func SetJoystickPlayerIndex(joystick *Joystick, player_index int32) bool {
//	return sdlSetJoystickPlayerIndex(joystick, player_index)
// }

// func SetJoystickVirtualAxis(joystick *Joystick, axis int32, value int16) bool {
//	return sdlSetJoystickVirtualAxis(joystick, axis, value)
// }

// func SetJoystickVirtualBall(joystick *Joystick, ball int32, xrel int16, yrel int16) bool {
//	return sdlSetJoystickVirtualBall(joystick, ball, xrel, yrel)
// }

// func SetJoystickVirtualButton(joystick *Joystick, button int32, down bool) bool {
//	return sdlSetJoystickVirtualButton(joystick, button, down)
// }

// func SetJoystickVirtualHat(joystick *Joystick, hat int32, value uint8) bool {
//	return sdlSetJoystickVirtualHat(joystick, hat, value)
// }

// func SetJoystickVirtualTouchpad(joystick *Joystick, touchpad int32, finger int32, down bool, x float32, y float32, pressure float32) bool {
//	return sdlSetJoystickVirtualTouchpad(joystick, touchpad, finger, down, x, y, pressure)
// }

// func UnlockJoysticks()  {
//	sdlUnlockJoysticks()
// }

// func UpdateJoysticks()  {
//	sdlUpdateJoysticks()
// }
