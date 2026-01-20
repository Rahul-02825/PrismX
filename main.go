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

// import (
// 	"io"
// 	"log"
// 	"net/http"
// 	"fmt"
// )

// func main() {
// 	h1 := func(w http.ResponseWriter, _ *http.Request) {
// 		io.WriteString(w, "Hello from a HandleFunc #1!\n")
// 	}
// 	h2 := func(w http.ResponseWriter, _ *http.Request) {
// 		io.WriteString(w, "Hello from a HandleFunc #2!\n")
// 	}

// 	http.HandleFunc("/", h1)
// 	http.HandleFunc("/endpoint", h2)
// 	fmt.Println("server is running")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }