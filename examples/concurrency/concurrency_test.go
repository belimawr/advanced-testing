package concurrencry

import (
	"fmt"
	"testing"
	"time"
)

//START OMIT
func Test_Slow1(t *testing.T) {
	t.Parallel() // HL
	time.Sleep(1 * time.Second)
	t.Log("Slow 1")
}

func Test_Slow2(t *testing.T) {
	t.Parallel() // HL
	time.Sleep(1 * time.Second)
	t.Log("Slow 2")
}

//END OMIT

//START-TABLE OMIT
func Test_table_test(t *testing.T) {
	t.Parallel()
	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("test-%d", i+1)
		t.Run(name, func(t *testing.T) {
			t.Parallel() // HL
			time.Sleep(time.Second)
		})
	}
}

//END-TABLE OMIT
