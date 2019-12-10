package concurrencry

import (
	"fmt"
	"testing"
	"time"
)

func Test_Slow1(t *testing.T) {
	time.Sleep(2 * time.Second)
	t.Log("Slow 1")
}

func Test_Slow2(t *testing.T) {
	time.Sleep(2 * time.Second)
	t.Log("Slow 2")
}

func Test_table_test(t *testing.T) {
	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("test-%d", i+1)
		t.Run(name, func(t *testing.T) {
			time.Sleep(time.Second)
		})
	}
}
