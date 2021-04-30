package server

import (
	"errors"
	"fmt"
	"net"
	"os"

	consul "github.com/hashicorp/consul/api"
	"github.com/rs/zerolog/log"
)

func (s *Server) register() error {
	if s.listener == nil {
		return errors.New("listener is null")
	}
	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return err
	}

	tcpAddr := s.listener.Addr().(*net.TCPAddr)
	port := tcpAddr.Port
	address, err := os.Hostname()
	if err != nil {
		address = "127.0.0.1"
	}
	log.Debug().Int("port", port).Str("addr", address).Msg("registering service...")
	agent := c.Agent()
	serviceDef := &consul.AgentServiceRegistration{
		Name:    s.conf.ServiceName,
		Port:    port,
		Address: address,
		Check: &consul.AgentServiceCheck{
			Interval: "10s",
			Timeout:  "3s",
			HTTP:     fmt.Sprintf("http://%s:%v/v2/healthz", address, port),
		},
	}

	if err := agent.ServiceRegister(serviceDef); err != nil {
		return err
	}
	s.agent = agent

	return nil
}

func (s *Server) deregister() {
	if s.agent != nil {
		s.agent.ServiceDeregister(s.conf.ServiceName)
	}
}
