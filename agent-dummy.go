// +build !amd64

package main

import (
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
	log.Fatal().Msg("agent not supported on this architecture")
	return nil
}

// Start the agent
func (a *Agent) Start() {}
