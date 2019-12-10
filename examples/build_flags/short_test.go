package buildflags

import (
	"testing"
	"time"
)

func Test_long(t *testing.T) {
	if testing.Short() { // HL
		t.Skip("skipping long test") // HL
	} // HL

	for i := 0; i < 10; i++ {
		t.Log("counting: ", i+1)
		time.Sleep(time.Second)
	}
}
