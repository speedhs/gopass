// components/connectDb.go

package components

import (
    "context"
    "fmt"
    "os"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// ConnectDb initializes and connects to the MongoDB database.
func ConnectDb() {
    // Use the SetServerAPIOptions() method to set the Stable API version to 1
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    errenv := godotenv.Load(".env")
    if errenv != nil {
        fmt.Println("Error loading .env file:", errenv)
        return
    }
    uri := os.Getenv("MONGODB_URI")
    opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

    // Create a new client and connect to the server
    ctx := context.TODO()
    var err error
    client, err = mongo.Connect(ctx, opts)
    if err != nil {
        fmt.Println("Error connecting to MongoDB:", err)
        return
    }
    // Ensure a successful connection by sending a ping
    if err := client.Ping(ctx, nil); err != nil {
        fmt.Println("Error pinging MongoDB:", err)
        return
    }
    // Printing for testing
    fmt.Println("You successfully connected to MongoDB!")
}
