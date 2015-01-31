package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const DELAY = 25

func cleanExit(sigVal os.Signal) {
	fmt.Printf("Exiting with signal: %v\n", sigVal)
	os.Exit(1)
}

func main() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	signal.Notify(sigChan, syscall.SIGKILL)

	go func() {
		sigVal := <-sigChan
		cleanExit(sigVal)
	}()

	for {
		time.Sleep(DELAY * time.Second)
		fmt.Printf("Hello World!\n")
	}
}
