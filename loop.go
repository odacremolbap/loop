package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	flag "github.com/spf13/pflag"
)

var iterations int
var exitStatus int

func init() {
	flag.IntVarP(&iterations, "iterations", "i", 10, "loop iterations")
	flag.IntVarP(&exitStatus, "exit", "e", 0, "exit status")
	flag.Parse()
}

func main() {

	// register for signals
	s := make(chan os.Signal, 1)
	signal.Notify(s,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR1,
		syscall.SIGUSR2)

	// react to signals
	go func() {
		received := <-s
		fmt.Printf(">>>>>>>>> Signal received: %s\n", received.String())
	}()

	// loop
	for i := 1; i <= iterations; i++ {
		fmt.Printf("Iteration #%d\n", i)
		time.Sleep(time.Second)
	}

	os.Exit(exitStatus)
}
