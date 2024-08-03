package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Ricid Kumbara"
}

func OnlyIn(channel chan<- string) {
	channel <- "Ricid Kumbara"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestCreateChannel(t *testing.T) {
	// channel := make(chan string)
	// channel <- "Ricid"
	// data := <-channel

	// fmt.Println(data)
	// defer close(channel)

	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Ricid Kumbara X"
		fmt.Println("Selesai Mengirim Data")
	}()

	fmt.Println("Menunggu Data")
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	// channel <- "Ricid"
	// channel <- "Kumbara"
	// channel <- "Fulan"

	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	go func() {
		channel <- "Ricid"
		channel <- "Kumbara"
		channel <- "Fulan"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan -" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima Data", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// select {
	// case data := <-channel1:
	// 	fmt.Println("Data dari channel 1", data)
	// case data := <-channel2:
	// 	fmt.Println("Data dari channel 2", data)
	// }

	// select {
	// case data := <-channel1:
	// 	fmt.Println("Data dari channel 1", data)
	// case data := <-channel2:
	// 	fmt.Println("Data dari channel 2", data)
	// }

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data", time.Now().UnixNano())
		}

		if counter == 2 {
			break
		}
	}
}

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)
}
