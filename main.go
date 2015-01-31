package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const DELAY = 25

func cleanExit() {
	fmt.Printf("Exiting...\n")
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	signal.Notify(sigChan, syscall.SIGKILL)
	go func() {
		<-sigChan
		cleanExit()
		os.Exit(1)
	}()

	for {
		time.Sleep(DELAY * time.Second)
		fmt.Printf("Hello World!\n")
	}
}
