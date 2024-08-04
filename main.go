package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	defer close(ch)

	ch <- "Sending Message"
	receive := <-ch

	fmt.Println(receive)
}
