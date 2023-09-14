package components

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)


func GetUserChoice() int {
	var choice int
	fmt.Print(Cyan+"Enter your choice: "+Reset)
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(Red+"Invalid input. Please enter a valid number."+Reset)
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