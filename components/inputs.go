package components

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

func GetUserChoice() int {
	var choice int
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return -1
	}
	return choice
}

func GetUserInfo(prompt string) string {
    fmt.Print(prompt)
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    return strings.TrimSpace(input)
}