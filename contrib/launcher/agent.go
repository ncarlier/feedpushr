// +build amd64

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"strings"

	"github.com/getlantern/systray"
	"github.com/ncarlier/feedpushr/v3/pkg/assets"
)

type fn func()

// Agent manage the service with the systray
type Agent struct {
	onReady fn
	onExit  fn
}

// NewAgent creates a new service agent
func NewAgent(onStart, onStop fn, url string) (*Agent, error) {
	iconPath := "/ui/logo.png"
	if runtime.GOOS == "windows" {
		iconPath = "/ui/favicon.ico"
	}
	file, err := assets.GetFS().Open(iconPath)
	if err != nil {
		return nil, fmt.Errorf("unable to load agent icon: %v", err)
	}
	icon, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to load agent icon: %v", err)
	}
	onReady := func() {
		systray.SetIcon(icon)
		systray.SetTitle("feedpushr")
		systray.SetTooltip("Open feedpushr menu")
		mOpen := systray.AddMenuItem("Open web UI", "Open UI in your browser")
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit", "Shutdown feedpushr")
		go func() {
			for {
				select {
				case <-mOpen.ClickedCh:
					err := openbrowser(url)
					if err != nil {
						log.Println("unable to open web UI", err)
					}
				case <-mQuit.ClickedCh:
					systray.Quit()
				}
			}
		}()
		onStart()
	}

	return &Agent{
		onReady: onReady,
		onExit:  onStop,
	}, nil
}

// Start the agent
func (a *Agent) Start() {
	systray.Run(a.onReady, a.onExit)
}

func openbrowser(url string) error {
	if !strings.HasPrefix(url, "http") {
		if strings.HasPrefix(url, ":") {
			url = "localhost" + url
		}
		url = "http://" + url
	}
	url = url + "/ui"
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}
