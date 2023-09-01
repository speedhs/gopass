package components

import(
	"fmt"
	
)

func AddInfo(){
	platform := GetUserInfo("Enter the platform name: ")
    username := GetUserInfo("Enter the username: ")
    password := GetUserInfo("Enter the password: ")

    // Here you can store or process the platform, username, and password as needed
    // For example, you can save them to a file, store them in a database, or perform encryption

    // Example: Print the entered data
    fmt.Println("You entered the following information:")
    fmt.Printf("Platform: %s\n", platform)
    fmt.Printf("Username: %s\n", username)
    fmt.Printf("Password: %s\n", password)
}