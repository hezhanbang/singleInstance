package main

import (
	"fmt"
	"time"

	"github.com/liyoubdu/singleInstance"
)

func main() {
	fmt.Println("start...")
	for i := 0; i < 9; i++ {
		time.Sleep(time.Second * 9)
		canRun, err := singleInstance.CurrentProcessIsSingle("__yourTest_2__", "")
		if err != nil {
			fmt.Printf("fail to check single, err=%s\n", err)
			time.Sleep(time.Hour)
			return
		}
		if canRun {
			fmt.Printf("current process is only single, can run now %v\n", time.Now().String())
		} else {
			fmt.Printf("too many processes are runing, exit %v\n", time.Now().String())
		}
	}
	fmt.Printf("done to check\n")
}
