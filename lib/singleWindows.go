package SingleInstance

import (
	"fmt"
	"syscall"
	"unsafe"
)

//HelloTest is fun
func HelloTest() {
	fmt.Printf("SingleInstance say hello to you!\n")
}

//IsSingle is fun
func IsSingle(key string) bool {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	procCreateMutex := kernel32.NewProc("CreateMutexW")

	_, _, err := procCreateMutex.Call(
		0,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(key))),
	)

	switch int(err.(syscall.Errno)) {
	case 0:
		return true
	default:
		return false
	}

	return true
}
