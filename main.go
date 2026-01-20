package main

import (
	"time"
	"PrismX/logger"
	"PrismX/loadBalancer"
	"PrismX/internal/database"

)

func main() {

	log := logger.InitLogger("app.log")

	defer log.Close()

	log.Info("Application started")

	log.Info("Running database connection")
	database.ConnectDatabase()

	go loadBalancer.StartLoadBalancer()

	for i := 0; i < 5; i++ {
		log.Info("Main thread working...")
		time.Sleep(1 * time.Second)
	}
	log.Info("Application exiting\n")
}
