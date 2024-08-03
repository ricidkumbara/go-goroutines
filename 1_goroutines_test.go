package main

import (
	"fmt"
	"testing"
	"time"
)

func RunSayHello() {
	fmt.Println(time.Now().Local(), "Hello, world!!!")
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestCreateGoroutines(t *testing.T) {
	go RunSayHello()
	fmt.Println(time.Now().Local(), "Done...")

	time.Sleep(1 * time.Second)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)
}
