// +build amd64

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"

	"github.com/getlantern/systray"
	"github.com/ncarlier/feedpushr/pkg/assets"
	"github.com/ncarlier/feedpushr/pkg/config"
	"github.com/rs/zerolog/log"
)

type fn func()

// Agent manage the service with the systray
type Agent struct {
	onReady fn
	onExit  fn
}

// NewAgent creates a new service agent
func NewAgent(onStart, onStop fn, conf config.Config) *Agent {
	iconPath := "/ui/logo.png"
	if runtime.GOOS == "windows" {
		iconPath = "/ui/favicon.ico"
	}
	file, err := assets.GetFS().Open(iconPath)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to load agent icon")
	}
	icon, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to load agent icon")
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
					err := openbrowser(conf.ListenAddr)
					if err != nil {
						log.Error().Err(err).Msg("unable to open web UI")
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
	}
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
