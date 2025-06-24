package util

import (
	"bufio"
	"fmt"
	"gohexarc/internal/domain"
	"os"
	"strings"
)

const List string = "list"
const Create string = "create"
const Update string = "update"
const Delete string = "delete"
const Get string = "get"

func PrintUsage() {
	fmt.Println("\nAvailable cli commands:")
	fmt.Println("list   - List users")
	fmt.Println("create - Create user")
	fmt.Println("update - Update user")
	fmt.Println("delete - Delete user")
	fmt.Println("get - Get user")
	fmt.Println("exit - Exit the CLI")
}

func PrintUser(user domain.User) {
	fmt.Println("\n-----User Details-----")
	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Println("----------------------")
}

func ReadInput(inputText string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n" + inputText)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
