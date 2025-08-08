package main

import (
	"fmt"
	"time"
)

func main() {

	var channelData = make(chan string)
	var channelClose = make(chan interface{})
	done := make(chan struct{})

	go func() {
		fruits := []string{"apple", "peach", "pear", "banana", "error"}

		for _, fruit := range fruits {
			if fruit == "error" {
				channelClose <- nil
				return
			}
			fmt.Println("Preparando un licuado de: ", fruit)
			channelData <- fruit
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case data := <-channelData:
				fmt.Println("Recibiendo un licuado de ", data)
			case data := <-channelClose:
				if data == nil {
					close(channelData)
				}
				close(done)
				return
			}
		}
	}()
	result := <-done
	fmt.Println("Fin del programa...", result)
}
