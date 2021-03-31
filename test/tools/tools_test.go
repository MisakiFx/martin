package tools

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/MisakiFx/martin/pkg/tools"
)

func TestSomeFunc(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	go func() {
		time.Sleep(5 * time.Second)
		close(ch)
	}()
	for i := range ch {
		fmt.Printf("%v", i)
	}
}

func TestGenId(t *testing.T) {
	log.Printf("gen id : %v", tools.GenId())
}
