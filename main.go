package main

import "fmt"

type data struct {
	platform string
	username string
	pass     string
}

func addpass(obj data) {

}
func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println("hi" + fmt.Sprint(x))
	switch x {
	case 1: //Add data
	case 2: //Retrieve data
	case 3: //Generate strong password
	}
}
