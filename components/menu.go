package components

import (
	"fmt"
)

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
func MainMenu() {
	prinLogo()
	showOptions()
}
