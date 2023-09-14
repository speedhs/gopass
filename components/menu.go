package components

import (
	"fmt"
)


func prinLogo() {
	fmt.Println(Purple+`
   ____   _      _____  
  / ___| | |    | ____| 
 | |  _  | |    | |__   
 | |_| | | |___ |  __|  
  \____| |_____||_|
	`+Reset)
}
func showOptions() {
	fmt.Println(Purple+"CLI Password Manager")
	fmt.Println(Purple+"1. Add Password"+Reset)
	fmt.Println(Purple+"2. Retrieve Password"+Reset)
	fmt.Println(Purple+"3. Update Password"+Reset)
	fmt.Println(Purple+"4. Delete Password"+Reset)
	fmt.Println(Purple+"5. Generate Password"+Reset)
	fmt.Println(Purple+"6. Password Policy"+Reset)
	fmt.Println(Red+"7. Exit"+Reset)
}
func MainMenu() {
	prinLogo()
	showOptions()
}
