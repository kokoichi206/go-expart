package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	// func shout(ping chan string, pong chan string) {
	for {
		// From ping to pong
		// ping to variable s
		s, ok := <-ping
		if !ok {
			// do something
		}

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
		// fmt.Println("Executing loop")
	}
}

func main() {
	// Create 2 channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter q to quit)")

	for {
		fmt.Print("-> ")

		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}

		// send userinput to ping channel
		ping <- userInput

		// wait for a response
		response := <-pong
		fmt.Printf("Response: %s\n", response)
	}

	fmt.Println("All done. closing channels")
	close(ping)
	close(pong)
}
