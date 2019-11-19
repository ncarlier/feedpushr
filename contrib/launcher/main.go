package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func lookupArg(args []string, name string) (string, bool) {
	for idx, arg := range args {
		if (arg == name || arg == "-"+name) && idx+1 < len(args) {
			return args[idx+1], true
		}
	}
	return "", false
}

func main() {
	binary, err := exec.LookPath("feedpushr")
	if err != nil {
		binary = "./feedpushr"
	}
	args := os.Args[1:]

	cmd := exec.Command(binary, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("unable to pipe STDOUT for cmd", err)
	}

	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		log.Fatal("unable to start cmd", err)
	}

	waitChan := make(chan error, 1)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan)

	start := func() {
		go func() {
			waitChan <- cmd.Wait()
			close(waitChan)
		}()
	}

	stop := func() {
		cmd.Process.Signal(os.Interrupt)
	}

	// Lookup URL
	url := ":8080"
	if value, ok := lookupArg(args, "-addr"); ok {
		url = value
	} else if value, ok := os.LookupEnv("APP_ADDR"); ok {
		url = value
	}

	agent, err := NewAgent(start, stop, url)
	if err != nil {
		log.Fatal("unable to start agent", err)
	}
	go agent.Start()

	for {
		select {
		case sig := <-sigChan:
			// Forward signal
			if err := cmd.Process.Signal(sig); err != nil {
				log.Println("unable to forwrard signal", sig, err)
			}
		case err := <-waitChan:
			// Subprocess exited. Get the return code, if we can
			var waitStatus syscall.WaitStatus
			if exitError, ok := err.(*exec.ExitError); ok {
				waitStatus = exitError.Sys().(syscall.WaitStatus)
				os.Exit(waitStatus.ExitStatus())
			}
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
}
