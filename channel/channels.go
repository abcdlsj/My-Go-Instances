package main

import "fmt"

func main() {
	message := make(chan string)

	go func() { message <- "ping" }()
	go func() { message <- "ping2" }()

	msg := <-message
	msg2 := <-message

	fmt.Println(msg)
	fmt.Println(msg2)
}
