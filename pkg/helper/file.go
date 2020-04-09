package helper

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
)

// TailFile read file content until string sequence is met
func TailFile(filename string, until string) (chan string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	outputChan := make(chan string)

	go func() {
		defer watcher.Close()
		defer file.Close()
		defer close(outputChan)
		if err := watcher.Add(filename); err != nil {
			outputChan <- fmt.Sprintf("error: %s", err.Error())
			return
		}
		r := bufio.NewReader(file)
		for {
			by, err := r.ReadBytes('\n')
			if err != nil && err != io.EOF {
				outputChan <- fmt.Sprintf("error: %s", err.Error())
				return
			}
			if err == nil {
				data := strings.TrimSuffix(string(by), "\n")
				outputChan <- data
				if data == until {
					return
				}
				continue
			}
			if err = waitForChange(watcher); err != nil {
				outputChan <- fmt.Sprintf("error: %s", err.Error())
				return
			}
		}
	}()
	return outputChan, nil
}

func waitForChange(w *fsnotify.Watcher) error {
	for {
		select {
		case event := <-w.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				return nil
			}
		case err := <-w.Errors:
			return err
		}
	}
}
