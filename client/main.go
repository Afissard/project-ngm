package main

import (
	"fmt"
	"client/pigeon"
)

func main() {

	inChan := make(chan []string)
	outChan := make(chan []string)

	go pigeon.StartConnection("localhost:4000", inChan, outChan)

	inChan <- []string{"hello"}

	msg := <-outChan
	fmt.Printf("end %s\n", msg[0])
}
