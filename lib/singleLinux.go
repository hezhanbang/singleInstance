package singleInstance

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
	fmt.Printf("singleInstance say hello to you!\n")
}

//CurrentProcessIsSingle is fun
func CurrentProcessIsSingle(singleKey string) (singling bool, err error) {
	if len(singleKey) < 5 || len(singleKey) > 20 {
		return false, fmt.Errorf("invalid length of singleKey")
	}

	//open locker file
	exeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file, err := os.OpenFile(exeDir+"/pid.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return false, fmt.Errorf("can not open pid.txt file")
	}

	//check locked
	locked, newLocker := locked(file)

	if false == locked || false == newLocker {
		// fail to get locker,
		// or we have got locker early before this call 'locked(file)'.
		//
		// so we must to close locker file
		file.Close()
		file = nil

		if false == locked {
			singling = false
			err = nil
			return
		}
		if false == newLocker {
			singling = true
			err = nil
			return
		}
	}

	//we get NEW locker, update time to file
	file.Truncate(0)
	data := fmt.Sprintf("[%s] [pid=%d]\n", time.Now().String(), os.Getpid())
	file.WriteString(data)

	//do not close locker file

	singling = true
	err = nil
	return
}

func locked(file *os.File) (locked, newLocker bool) {
	if false == lockedByThis {
		err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
		if nil == err {
			//get new locker
			lockedByThis = true
			locked = true
			newLocker = true
		} else {
			////fail to get locker
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
