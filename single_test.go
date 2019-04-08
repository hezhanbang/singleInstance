package SingleInstance

import "testing"

func TestRun(t *testing.T) {
	HelloTest()
	IsSingle()

	t.Logf("---------DONE------")
}
