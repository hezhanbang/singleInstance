package SingleInstance

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"
)

var lockedByThis = false

//HelloTest is fun
func HelloTest() {
	fmt.Printf("SingleInstance say hello to you!\n")
}

//CurrentProcessIsSingle is fun
func CurrentProcessIsSingle(singleKey string) (singling, ok bool) {
	if len(singleKey) < 5 || len(singleKey) > 20 {
		return false, false
	}
	locked, newLocker := locked(singleKey)
	if !locked {
		return false, true
	}
	if !newLocker {
		return true, true
	}

	//we get new locker, update time to file
	exeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file, err := os.OpenFile(exeDir+"\\pid.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return false, false
	}

	data := fmt.Sprintf("[%s] [pid=%d]\n", time.Now().String(), os.Getpid())
	n, err := file.WriteString(data)
	if err != nil || n != len(data) {
		return true, false
	}

	file.Close()
	file = nil
	return true, true
}

func locked(key string) (locked, newLocker bool) {
	if false == lockedByThis {
		//test for new locker
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		procCreateMutex := kernel32.NewProc("CreateMutexW")
		closeHandle := kernel32.NewProc("CloseHandle")

		//call CreateMutex
		handle, _, err := procCreateMutex.Call(
			0,
			1,
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(key))),
		)

		fmt.Printf("CreateMutexW, handle=%d errInt=%d errStr=%v [SingleInstance]\n", handle, int(err.(syscall.Errno)), err)

		//check return val and last err
		if 0 == int(err.(syscall.Errno)) {
			lockedByThis = true
			locked = true
			newLocker = true
		} else { //fail to get locker, we have to release reference count of the kernel object.
			if handle != 0 {
				closeHandle.Call(handle)
			}
			lockedByThis = false
			locked = false
			newLocker = false
		}

		return
	}

	//we have keep this locker
	locked = true
	newLocker = false
	return
}
