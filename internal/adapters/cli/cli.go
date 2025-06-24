package cli

import (
	"fmt"
	"gohexarc/internal/domain"
	"gohexarc/internal/port"
	"os"
)

const list string = "list"
const create string = "create"
const update string = "update"
const delete string = "delete"
const get string = "get"

func RunCli(service port.UserService) {
	if len(os.Args) == 2 {
		printUsage()
		return
	}

	command := os.Args[2]
	switch command {
	case list:
		listUsers(service)
	case get:
		getUser(service)
	case create:
		createUser(service)
	case update:
		updateUser(service)
	case delete:
		deleteUser(service)
	}
}

func listUsers(service port.UserService) {
	users, err := service.ListUsers()
	if err != nil {
		err = fmt.Errorf("could not list users: %v", err)
		return
	}
	for _, user := range users {
		printUser(user)
	}
}

func getUser(service port.UserService) {
	user, err := service.GetUser(os.Args[3])
	if err != nil {
		errorMessage := fmt.Errorf("user %q not found", os.Args[3])
		fmt.Println(errorMessage)
		return
	}
	printUser(user)
}

func createUser(service port.UserService) {
	user, err := service.CreateUser(os.Args[3], os.Args[4])
	if err != nil {
		errorMessage := fmt.Errorf("could not create user %q - %q", os.Args[3], os.Args[4])
		fmt.Println(errorMessage)
	}
	fmt.Println("User created successfully")
	printUser(user)
}

func updateUser(service port.UserService) {
	err := service.UpdateUser(os.Args[3], os.Args[4], os.Args[5])
	if err != nil {
		errorMessage := fmt.Errorf("could not update user %q - %q - %q", os.Args[3], os.Args[4], os.Args[5])
		fmt.Println(errorMessage)
	}
	fmt.Println("User updated successfully")
	user := domain.User{ID: os.Args[3], Name: os.Args[4], Email: os.Args[5]}
	printUser(user)
}

func deleteUser(service port.UserService) {
	err := service.DeleteUser(os.Args[3])
	if err != nil {
		errorMessage := fmt.Errorf("could not delete user %q", os.Args[3])
		fmt.Println(errorMessage)
	}
	fmt.Println("User deleted successfully")
}

func printUsage() {
	fmt.Println("Available cli commands:")
	fmt.Println("  list   List users")
	fmt.Println("  create Create user")
	fmt.Println("  update Update user")
	fmt.Println("  delete Delete user")
}

func printUser(user domain.User) {
	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Email: %s\n", user.Email)
}
