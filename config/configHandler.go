package config

// configs has to be added later via db
type config struct {
	servers []string
	balancer string
	replicas int
}

func LoadConfig() (*config, error) {
	return &config{
		servers: []string{"server1", "server2","server3"},
	}, nil
}

func (c *config)GetServers() []string{
	return c.servers
}
