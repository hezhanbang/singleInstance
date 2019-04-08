package SingleInstance

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	//HelloTest()

	for i := 0; i < 5; i++ {
		ok := IsSingle("__yourTest___")
		if ok {
			fmt.Printf("current process is only single, can run now\n")
		} else {
			fmt.Printf("too many processes are runing, exit now\n")
		}
	}
	fmt.Printf("done to check\n")

	t.Logf("---------DONE------")
}
