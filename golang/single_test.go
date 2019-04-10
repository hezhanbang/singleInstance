package singleInstance

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	//HelloTest()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 3)
		single, err := CurrentProcessIsSingle("__yourTest___", "")
		if err != nil {
			fmt.Printf("fail to do CurrentProcessIsSingle, err=%s\n", err)
			return
		}

		if single {
			fmt.Printf("current process is only single, can run now %v\n", time.Now().String())
		} else {
			fmt.Printf("too many processes are runing, exit %v\n", time.Now().String())
		}
	}
	fmt.Printf("done to check\n")

	t.Logf("---------DONE------")
}
