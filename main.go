package main

import (
	"time"

	"PrismX/logger"
	"PrismX/loadBalancer"
	"PrismX/internal/database"
	"PrismX/internal/controller"
	"net/http"


)

func main() {

	log := logger.InitLogger("app.log")

	defer log.Close()

	log.Info("Application started")

	log.Info("Running database connection")
	database.ConnectDatabase()

	go loadBalancer.StartLoadBalancer()

	http.HandleFunc("/createuser",controller.CreateUser)
	

	for i := 0; i < 5; i++ {
		log.Info("Main thread working...")
		time.Sleep(1 * time.Second)
	}

	log.Info("Proxy server is running on port :8080\n")
	http.ListenAndServe(":8080", nil)
}
