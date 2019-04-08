package SingleInstance

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

var gFlag = os.O_CREATE | os.O_RDWR | os.O_TRUNC
var lockedByThis = false

//HelloTest is fun
func HelloTest() {
	fmt.Printf("SingleInstance say hello to you!\n")
}

//CurrentProcessIsSingle is fun
func CurrentProcessIsSingle(singleKey string) bool {
	exeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	file, err := os.OpenFile(exeDir+"/pid.txt", gFlag, 0666)
	if err != nil {
		return false
	}

	gFlag = os.O_CREATE | os.O_RDWR

	if false == lockedByThis {
		err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
		if err != nil {
			file.Close()
			file = nil

			lockedByThis = false
			return false //return
		}
		lockedByThis = true
	}

	file.Truncate(0)
	data := fmt.Sprintf("[%s] [pid=%d]\n", time.Now().String(), os.Getpid())
	file.WriteString(data)
	return true
}
