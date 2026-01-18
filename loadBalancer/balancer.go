package loadBalancer


// var log = logger.InitLogger("app.log")

// factory design pattern implementation for the balancers
type loadbalancer interface{
	insertServer(server string)
	removeServer(server string)	
	getServer(request string) string
}
// for now consistent hashing is only available balancer
func balancerFactory(balancerType string) (loadbalancer){

	switch balancerType{
	case "consistent-hash":
			return &ConsistentHash{}
		
	default:
		return nil
	}	
	
}

