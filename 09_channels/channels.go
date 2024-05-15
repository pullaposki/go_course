package main

import (
	"fmt"
	"time"
)

func doWork(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second*3)
	fmt.Println("done")

	fmt.Println("Worker: Informing main that we are done")
	done <- true
}

func main() {
	messages := make (chan string)

	fmt.Println("-- Basic channel --")

	// Start a goroutine to send a message to the channel
	go func () {
		fmt.Println(("Pinger: pining the main"))
		messages <- "ping"
	}()

	fmt.Println("Main: reading the channel")

	// Read the message from the channel
	msg := <- messages

	// Print the message
	fmt.Println((msg))

	time.Sleep(2*time.Second)

	fmt.Println("-- Buffered channel --")

	// Create a buffered channel with a capacity of 2
	bufferedMessages := make (chan string, 2)

	// Send two messages to the channel
	bufferedMessages <- "buffered"
	bufferedMessages <- "channel"

	// Read the messages from the channel
	fmt.Println(<-bufferedMessages)
	fmt.Println(<-bufferedMessages)

	time.Sleep(2*time.Second)

	fmt.Println("-- Channel synchronization --")

	// Create a channel to notify the main goroutine that the worker is done
	done := make(chan bool)

	// Start the worker in a goroutine
	go doWork(done)
	
	fmt.Println("Main: waiting for the worker to finish")
	
	// Block until we receive a notification from the worker
	<-done

	fmt.Println("Main: worker done. Exiting...")

}