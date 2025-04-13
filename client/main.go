package main

import (
	"fmt"
)

func main() {

	inChan := make(chan []string)
	outChan := make(chan []string)

	pigeon.StartConnection("localhost:4000", inChan, outChan)

	inChan <- []string{"hello"}

	msg := <-outChan
	fmt.Printf("Received: %s\n", msg[0])
}
