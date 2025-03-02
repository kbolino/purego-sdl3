package sdl

type AsyncIOTaskType uint32

const (
	AsyncIOTaskRead AsyncIOTaskType = iota
	AsyncIOTaskWrite
	AsyncIOTaskClose
)

type AsyncIOResult uint32

const (
	AsyncIOComplete AsyncIOResult = iota
	AsyncIOFailure
	AsyncIOCanceled
)

// func AsyncIOFromFile(file string, mode string) *AsyncIO {
//	return sdlAsyncIOFromFile(file, mode)
// }

// func CloseAsyncIO(asyncio *AsyncIO, flush bool, queue *AsyncIOQueue, userdata unsafe.Pointer) bool {
//	return sdlCloseAsyncIO(asyncio, flush, queue, userdata)
// }

// func CreateAsyncIOQueue() *AsyncIOQueue {
//	return sdlCreateAsyncIOQueue()
// }

// func DestroyAsyncIOQueue(queue *AsyncIOQueue)  {
//	sdlDestroyAsyncIOQueue(queue)
// }

// func GetAsyncIOResult(queue *AsyncIOQueue, outcome *AsyncIOOutcome) bool {
//	return sdlGetAsyncIOResult(queue, outcome)
// }

// func GetAsyncIOSize(asyncio *AsyncIO) int64 {
//	return sdlGetAsyncIOSize(asyncio)
// }

// func LoadFileAsync(file string, queue *AsyncIOQueue, userdata unsafe.Pointer) bool {
//	return sdlLoadFileAsync(file, queue, userdata)
// }

// func ReadAsyncIO(asyncio *AsyncIO, ptr unsafe.Pointer, offset uint64, size uint64, queue *AsyncIOQueue, userdata unsafe.Pointer) bool {
//	return sdlReadAsyncIO(asyncio, ptr, offset, size, queue, userdata)
// }

// func SignalAsyncIOQueue(queue *AsyncIOQueue)  {
//	sdlSignalAsyncIOQueue(queue)
// }

// func WaitAsyncIOResult(queue *AsyncIOQueue, outcome *AsyncIOOutcome, timeoutMS int32) bool {
//	return sdlWaitAsyncIOResult(queue, outcome, timeoutMS)
// }

// func WriteAsyncIO(asyncio *AsyncIO, ptr unsafe.Pointer, offset uint64, size uint64, queue *AsyncIOQueue, userdata unsafe.Pointer) bool {
//	return sdlWriteAsyncIO(asyncio, ptr, offset, size, queue, userdata)
// }
