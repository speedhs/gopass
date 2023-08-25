package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type data struct {
	platform string
	username string
	pass     string
}

func getUserChoice() int {
	var choice int
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return -1
	}
	return choice
}
func prinLogo() {
	fmt.Println(`
   ____   _      _____  
  / ___| | |    | ____| 
 | |  _  | |    | |__   
 | |_| | | |___ |  __|  
  \____| |_____||_|
	`)
}
func showOptions() {
	fmt.Println("CLI Password Manager")
	fmt.Println("1. Add Password")
	fmt.Println("2. Retrieve Password")
	fmt.Println("3. Update Password")
	fmt.Println("4. Delete Password")
	fmt.Println("5. Generate Password")
	fmt.Println("6. Password Policy")
	fmt.Println("7. Exit")
}
func mainMenu() {
	prinLogo()
	showOptions()
}
func addpass(obj data) {

}

func connectDb() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	uri := os.Getenv("MONGODB_URI")
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	//Printing for testing
	fmt.Println("You successfully connected to MongoDB!")

}
func main() {
	connectDb()
	mainMenu()
	for true {
		choice := getUserChoice()
		switch choice {
		case 1:
			fmt.Println("Adding a new password...")
		case 2:
			fmt.Println("Retrieving a password...")
		case 3:
			fmt.Println("Updating a password...")
		case 4:
			fmt.Println("Deleting a password...")
		case 5:
			fmt.Println("Generating password...")
		case 6:
			fmt.Println("Password Policy")
		case 7:
			fmt.Println("Exiting the password manager...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
