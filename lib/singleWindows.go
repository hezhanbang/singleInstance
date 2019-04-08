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
func CurrentProcessIsSingle(singleKey string) bool {
	locked, newLocker := locked(singleKey)
	if !locked {
		return false
	}
	if !newLocker {
		return true
	}

	//we get new locker, update time to file
	exeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file, err := os.OpenFile(exeDir+"\\pid.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return false
	}

	data := fmt.Sprintf("[%s] [pid=%d]\n", time.Now().String(), os.Getpid())
	file.WriteString(data)

	file.Close()
	file = nil
	return true
}

func locked(key string) (locked, newLocker bool) {
	if false == lockedByThis {
		//test for new locker
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		procCreateMutex := kernel32.NewProc("CreateMutexW")

		_, _, err := procCreateMutex.Call(
			0,
			0,
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(key))),
		)

		if 0 == int(err.(syscall.Errno)) {
			lockedByThis = true
			locked = true
			newLocker = true
		} else {
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
