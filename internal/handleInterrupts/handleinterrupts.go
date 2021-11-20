package handleInterrupts

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var WindowsInterrupt = make(chan bool, 1)

func HandleStop() {
	go handleStopWindows()
	channel := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-channel // This blocks the main thread until an interrupt is received

	log.Println("Paragüero was successful shutdown.")
	os.Exit(1)
}

// handleStopWindows handles ctrl c situations on Windows that only return an EOF instead of a signal term
func handleStopWindows() {
	if runtime.GOOS == "windows" {
		_ = <-WindowsInterrupt
		log.Println("Paragüero was successful shutdown.")
		os.Exit(1)
	} else {
		WindowsInterrupt = nil
	}
}
