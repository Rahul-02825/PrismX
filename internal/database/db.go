package database


import (
	"context"
	"os"
	"PrismX/logger"
	// "go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"github.com/joho/godotenv"	

)


var (
	Client *mongo.Client
	UserCollection *mongo.Collection
	ConfigCollection *mongo.Collection
)

func ConnectDatabase() {
	log := logger.InitLogger("app.log")

	var uri string
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}
	if uri = os.Getenv("MONGO_URL"); uri == "" {
		log.Error("Incorrect Mongodb url or unknown url")
	}

	// Uses the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Defines the options for the MongoDB client
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Creates a new client and connects to the server
	client, err := mongo.Connect(opts)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	log.Info("Successfully connected to Database")
	Client = client

	db := client.Database("PrismX")
	UserCollection := db.Collection("User")
	ConfigCollection := db.Collection("Config")

	
	// fmt.Println("connection to databse is going on")
	// // Sends a ping to confirm a successful connection
	// var result bson.M
	// if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
	// 	fmt.Println("some error occured")
	// 	panic(err)
	// }
	
}

func GetDatabaseClient() *mongo.Client{
	return Client
}