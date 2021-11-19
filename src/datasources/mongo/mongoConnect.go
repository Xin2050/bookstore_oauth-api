package mongo_con

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

const (
	mongo_username = "mongo_username"
	mongo_password = "mongo_password"
	mongo_host     = "mongo_host"
	mongo_db       = "mongo_db"
)

var (
	client   *mongo.Client
	Database string
)

func GetClient() *mongo.Client {
	if client == nil {
		initClient()
	}
	return client
}
func GetDatabase() *mongo.Database {
	if client == nil {
		initClient()
	}
	return client.Database(Database)
}
func initClient() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv(mongo_username)
	password := os.Getenv(mongo_password)
	host := os.Getenv(mongo_host)
	Database = os.Getenv(mongo_db)
	connectUri := fmt.Sprintf("mongodb://%s:%s@%s/?maxPoolSize=20&w=majority", username, password, host)
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(connectUri))
	if err != nil {
		log.Fatal("connect error:", err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("connect error:", err)
	}

	fmt.Println("Successfully connected and pinged.")
}
