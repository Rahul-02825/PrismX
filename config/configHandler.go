package config

// configs has to be added later via db
type upstream struct {
	name     string
	lbMethod string
	servers  []upstreamservers
	replicas int
}

type config struct {
	upstream upstream
	balancer  string
	replicas  int
	server []server
}
type upstreamservers struct{
	address string
	weight int
	maxFails int
	FailTimeout int
	down bool
}

type server struct{
	serverName string
}


func LoadConfig() (*config, error) {
	return &config{
		upstream: upstream{
				name:     "Auth service",
				lbMethod: "consistent-hash",
				replicas: 3,
				servers:  []upstreamservers{
					{
						address:"http://localhost:9000/auth",
						weight:10,
						maxFails: 2,
						FailTimeout: 2,
						down:false,
					},
				},
			},
		server:[]server{
			{serverName:"RetailService"},
		},
	}, nil
}


func (c *config) GetServers() upstream {
	return c.upstream
}
