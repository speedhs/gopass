package main

import (
	
	"fmt"
	"os"
	
	com "gopass/components"
)



func main() {
	com.ConnectDb()
	com.MainMenu()
	for true {
		choice := com.GetUserChoice()
		switch choice {
		case 1:
			fmt.Println("Adding a new password...")
			com.AddInfo()
		case 2:
			fmt.Println("Retrieving a password...")
			com.PrintUsernameAndPassword()
		case 3:
			fmt.Println("Updating a password...")
			com.UpdatePassword()
		case 4:
			fmt.Println("Deleting a password...")
			com.DeletePassword()
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
