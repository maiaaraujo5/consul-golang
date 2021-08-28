package consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/connect"
	"log"
	"time"
)

type Consul struct {
	Agent *api.Agent
	TTl   time.Duration
}

func Connect() (*connect.Service, error) {
	consul := &Consul{
		TTl: time.Minute,
	}

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}

	consul.Agent = client.Agent()

	service := &api.AgentServiceRegistration{
		Name: "consul-golang",
		Check: &api.AgentServiceCheck{
			TTL: consul.TTl.String(),
		},
	}

	err = consul.Agent.ServiceRegister(service)
	if err != nil {
		return nil, err
	}

	go consul.HeartBeat()

	return nil, nil
}

func (c *Consul) HeartBeat() {
	ticker := time.NewTicker(c.TTl / 2)
	for range ticker.C {
		err := c.Agent.PassTTL("service:consul-golang", "")
		if err != nil {
			log.Println(err)
		}
	}
}
