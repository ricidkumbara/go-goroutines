package main

func main() {
	ch := make(chan string, 3)
	defer close(ch)
}
