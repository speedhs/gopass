// components/functions.go

package components

import (
    "fmt"
    "context"
    "go.mongodb.org/mongo-driver/bson" // Import the bson package
    "time" // Import the time package
    // "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
    // "go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo"
)

// Book represents a book entity.
type Data struct {
	Id string `json:"id"`
    Platform string `json:"author"`
    Username  string `json:"title"`
    Password string `json:password`
	// Add other fields as needed, such as publication year, genre, etc.
}

var (
    Reset   = "\033[0m"
    Red     = "\033[31m"
    Green   = "\033[32m"
    Yellow  = "\033[33m"
    Blue    = "\033[34m"
    Purple  = "\033[35m"
    Cyan    = "\033[36m"
    White   = "\033[97m"
)

// AddInfo adds information to the MongoDB database.
func AddInfo() {
    if client == nil {
        fmt.Println("MongoDB client is nil")
        return
    }
	id:="1"
	platform := GetUserInfo(Blue+"Enter the platform name: "+Reset)
    username := GetUserInfo(Blue+"Enter the username: "+Reset)
    password := GetUserInfo(Blue+"Enter the password: "+Reset)
    // Here you can store or process the platform, username, and password as needed
    // For example, you can save them to a file, store them in a database, or perform encryption
    // Example: Print the entered data
    // fmt.Println("You entered the following information:")
    // fmt.Printf("Platform: %s\n", platform)
    // fmt.Printf("Username: %s\n", username)
    // fmt.Printf("Password: %s\n", password)
    coll := client.Database("dummy").Collection((id))
    doc := Data{Id: id, Platform: platform, Username: username, Password: password}

    result, err := coll.InsertOne(context.TODO(), doc)
    if err != nil {
        fmt.Println("Error inserting document:", err)
        return
    }
    fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
		
}

// PrintUsernameAndPassword prints the username and password for entries with platform 'test1'.
func PrintUsernameAndPassword() {
    if client == nil {
        fmt.Println("MongoDB client is nil")
        return
    }
    
    // platformToSearch := "test1"
    platformToSearch := GetUserInfo(Blue+"Enter the platform name: "+Reset)
    // fmt.Println(platformToSearch)
    // Create a filter to match documents with the specified platform
    filter := bson.M{"platform": platformToSearch}

    // Create a context with a timeout (adjust the timeout as needed)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Specify the collection to search in (replace "dummy" with your actual collection name)
    coll := client.Database("dummy").Collection("1")

    // Perform a find operation to retrieve documents matching the filter
    cursor, err := coll.Find(ctx, filter)
    if err != nil {
        fmt.Println("Error finding documents:", err)
        return
    }
    defer cursor.Close(ctx)

    // Iterate through the cursor and print the username and password of matching entries
    for cursor.Next(ctx) {
        var document Data
        if err := cursor.Decode(&document); err != nil {
            fmt.Println("Error decoding document:", err)
            return
        }
        fmt.Printf(Blue+"Username:" +Reset+ "%s\n", document.Username)
        fmt.Printf(Blue+"Password:" +Reset+ "%s\n", document.Password)
    }

    if err := cursor.Err(); err != nil {
        fmt.Println("Error iterating through documents:", err)
        return
    }
}
