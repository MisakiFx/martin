package tools

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/MisakiFx/martin/martin/pkg/tools"
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

func TestCheckTime(t *testing.T) {
	tools.Init()
	threeDaysLater := time.Now().Add(time.Hour * 24 * 3)
	limitTime := time.Date(threeDaysLater.Year(), threeDaysLater.Month(), threeDaysLater.Day(), 23, 59, 59, 0, tools.LocGloble)
	start := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+4, 0, 0, 0, 0, tools.LocGloble)
	if limitTime.Sub(start) < 0 {
		fmt.Printf("超出范围")
	} else {
		fmt.Printf("通过")
	}
}
