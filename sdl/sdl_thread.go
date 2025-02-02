package sdl

type ThreadPriority uint32

const (
	ThreadPriorityLow ThreadPriority = iota
	ThreadPriorityNormal
	ThreadPriorityHigh
	ThreadPriorityTimeCritical
)

type ThreadState uint32

const (
	ThreadUnknown ThreadState = iota
	ThreadAlive
	ThreadDetached
	ThreadComplete
)

// func CleanupTLS()  {
//	sdlCleanupTLS()
// }

// func CreateThreadRuntime(fn ThreadFunction, name string, data unsafe.Pointer, pfnBeginThread FunctionPointer, pfnEndThread FunctionPointer) *Thread {
//	return sdlCreateThreadRuntime(fn, name, data, pfnBeginThread, pfnEndThread)
// }

// func CreateThreadWithPropertiesRuntime(props PropertiesID, pfnBeginThread FunctionPointer, pfnEndThread FunctionPointer) *Thread {
//	return sdlCreateThreadWithPropertiesRuntime(props, pfnBeginThread, pfnEndThread)
// }

// func DetachThread(thread *Thread)  {
//	sdlDetachThread(thread)
// }

// func GetCurrentThreadID() ThreadID {
//	return sdlGetCurrentThreadID()
// }

// func GetThreadID(thread *Thread) ThreadID {
//	return sdlGetThreadID(thread)
// }

// func GetThreadName(thread *Thread) string {
//	return sdlGetThreadName(thread)
// }

// func GetThreadState(thread *Thread) ThreadState {
//	return sdlGetThreadState(thread)
// }

// func GetTLS(id *TLSID) unsafe.Pointer {
//	return sdlGetTLS(id)
// }

// func SetCurrentThreadPriority(priority ThreadPriority) bool {
//	return sdlSetCurrentThreadPriority(priority)
// }

// func SetTLS(id *TLSID, value unsafe.Pointer, destructor TLSDestructorCallback) bool {
//	return sdlSetTLS(id, value, destructor)
// }

// func WaitThread(thread *Thread, status *int32)  {
//	sdlWaitThread(thread, status)
// }
