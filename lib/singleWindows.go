package SingleInstance

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"
)

//HelloTest is fun
func HelloTest() {
	fmt.Printf("SingleInstance say hello to you!\n")
}

//CurrentProcessIsSingle is fun
func CurrentProcessIsSingle(singleKey string) bool {
	exeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	procCreateMutex := kernel32.NewProc("CreateMutexW")

	_, _, err := procCreateMutex.Call(
		0,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(singleKey))),
	)

	switch int(err.(syscall.Errno)) {
	case 0:
		return true
	default:
		return false
	}

	file, err := os.OpenFile(exeDir+"\\pid.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return false
	}

	data := fmt.Sprintf("[%s] [pid=%d]\n", time.Now().String(), os.Getpid())
	file.WriteString(data)
	return true
}
