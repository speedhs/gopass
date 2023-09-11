// components/functions.go

package components

import (
    "fmt"
    "context"
    // "go.mongodb.org/mongo-driver/mongo"
)

// Book represents a book entity.
type Book struct {
    Title  string `json:"title"`
    Author string `json:"author"`
    // Add other fields as needed, such as publication year, genre, etc.
}

// AddInfo adds information to the MongoDB database.
func AddInfo() {
    if client == nil {
        fmt.Println("MongoDB client is nil")
        return
    }
    coll := client.Database("Books").Collection("books")
    doc := Book{Title: "Atonent", Author: "Ian McEwan"}

    result, err := coll.InsertOne(context.TODO(), doc)
    if err != nil {
        fmt.Println("Error inserting document:", err)
        return
    }
    fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
    
    // Now you can continue with your user input and data processing logic
}
