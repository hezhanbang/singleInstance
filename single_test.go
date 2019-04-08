package SingleInstance

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	//HelloTest()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 3)
		ok := IsSingle("__yourTest___")
		if ok {
			fmt.Printf("current process is only single, can run now %v\n", time.Now().String())
		} else {
			fmt.Printf("too many processes are runing, exit %v\n", time.Now().String())
		}
	}
	fmt.Printf("done to check\n")

	t.Logf("---------DONE------")
}
