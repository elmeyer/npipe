// Code generated by 'go generate'; DO NOT EDIT.

package npipe

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procCancelIoEx          = modkernel32.NewProc("CancelIoEx")
	procConnectNamedPipe    = modkernel32.NewProc("ConnectNamedPipe")
	procCreateEventW        = modkernel32.NewProc("CreateEventW")
	procCreateNamedPipeW    = modkernel32.NewProc("CreateNamedPipeW")
	procDisconnectNamedPipe = modkernel32.NewProc("DisconnectNamedPipe")
	procGetOverlappedResult = modkernel32.NewProc("GetOverlappedResult")
	procWaitNamedPipeW      = modkernel32.NewProc("WaitNamedPipeW")
)

func cancelIoEx(handle syscall.Handle, overlapped *syscall.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall(procCancelIoEx.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func connectNamedPipe(handle syscall.Handle, overlapped *syscall.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall(procConnectNamedPipe.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func createEvent(sa *syscall.SecurityAttributes, manualReset bool, initialState bool, name *uint16) (handle syscall.Handle, err error) {
	var _p0 uint32
	if manualReset {
		_p0 = 1
	}
	var _p1 uint32
	if initialState {
		_p1 = 1
	}
	r0, _, e1 := syscall.Syscall6(procCreateEventW.Addr(), 4, uintptr(unsafe.Pointer(sa)), uintptr(_p0), uintptr(_p1), uintptr(unsafe.Pointer(name)), 0, 0)
	handle = syscall.Handle(r0)
	if handle == syscall.InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

func createNamedPipe(name *uint16, openMode uint32, pipeMode uint32, maxInstances uint32, outBufSize uint32, inBufSize uint32, defaultTimeout uint32, sa *syscall.SecurityAttributes) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall9(procCreateNamedPipeW.Addr(), 8, uintptr(unsafe.Pointer(name)), uintptr(openMode), uintptr(pipeMode), uintptr(maxInstances), uintptr(outBufSize), uintptr(inBufSize), uintptr(defaultTimeout), uintptr(unsafe.Pointer(sa)), 0)
	handle = syscall.Handle(r0)
	if handle == syscall.InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

func disconnectNamedPipe(handle syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDisconnectNamedPipe.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getOverlappedResult(handle syscall.Handle, overlapped *syscall.Overlapped, transferred *uint32, wait bool) (err error) {
	var _p0 uint32
	if wait {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall6(procGetOverlappedResult.Addr(), 4, uintptr(handle), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(transferred)), uintptr(_p0), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func waitNamedPipe(name *uint16, timeout uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procWaitNamedPipeW.Addr(), 2, uintptr(unsafe.Pointer(name)), uintptr(timeout), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}
