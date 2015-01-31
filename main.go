package main

import notify "github.com/mqu/go-notify"
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	DELAY     = 25
	NOTEDELAY = 10000
)

func sendNote(note *notify.NotifyNotification) {
	if note == nil {
		fmt.Printf("Note is nil\n")
		os.Exit(1)
	}

	note.SetTimeout(NOTEDELAY)
	if err := note.Show(); err != nil {
		// err is never nil? Catching SIGABRT from cgo execution.
		fmt.Printf("note.Show() caught error: %v\n", err.Message())
	}
	time.Sleep(NOTEDELAY * time.Millisecond)
	note.Close()
}

func cleanExit(sigVal os.Signal) {
	fmt.Printf("Exiting with signal: %v\n", sigVal)
	notify.UnInit()
	os.Exit(0)
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

	notify.Init("Nora")

	for {
		time.Sleep(DELAY * time.Second)
		note := notify.NotificationNew("Message from Nora", "Helpful text goes here.", "")
		go sendNote(note)
	}
}
