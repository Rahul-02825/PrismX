package main

import (
	"time"
	"PrismX/logger"
	"PrismX/loadBalancer"
)

func main() {
	log := logger.InitLogger("app.log")
	
	defer log.Close()

	log.Info("Application started")
	
	go loadBalancer.StartLoadbalancer()

	for i := 0; i < 5; i++ {
		log.Info("Main thread working...")
		time.Sleep(1 * time.Second)
	}
	log.Info("Application exiting\n")
}
