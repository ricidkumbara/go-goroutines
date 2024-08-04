package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsyc(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello, World!")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyc(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}
