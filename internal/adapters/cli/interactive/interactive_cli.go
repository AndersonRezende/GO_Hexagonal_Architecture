package interactiveCli

import (
	"fmt"
	"gohexarc/internal/adapters/cli/util"
	"gohexarc/internal/domain"
	"gohexarc/internal/port"
)

type Interactive struct {
	service port.UserService
}

func NewInteractiveCLI(service port.UserService) *Interactive {
	return &Interactive{service: service}
}

func (interactiveCli *Interactive) Run() {
	for {
		util.PrintUsage()
		command := util.ReadInput("Choose one of the available options above: ")
		if command == "exit" {
			fmt.Println("Exiting.")
			return
		}
		interactiveCli.ExecuteCommand(command)
	}
}

func (interactiveCli *Interactive) ExecuteCommand(command string) {
	switch command {
	case util.List:
		interactiveCli.ListUsers()
	case util.Get:
		interactiveCli.GetUser()
	case util.Create:
		interactiveCli.CreateUser()
	case util.Update:
		interactiveCli.UpdateUser()
	case util.Delete:
		interactiveCli.DeleteUser()
	}
}

func (interactiveCli *Interactive) ListUsers() {
	users, err := interactiveCli.service.ListUsers()
	if err != nil {
		fmt.Printf("could not list users: %v\n", err)
		return
	}
	for _, user := range users {
		util.PrintUser(user)
	}
}

func (interactiveCli *Interactive) GetUser() {
	id := util.ReadInput("Enter user ID: ")
	user, err := interactiveCli.service.GetUser(id)
	if err != nil {
		fmt.Printf("user %q not found\n", id)
		return
	}
	util.PrintUser(user)
}

func (interactiveCli *Interactive) CreateUser() {
	name := util.ReadInput("Enter name: ")
	email := util.ReadInput("Enter email: ")
	user, err := interactiveCli.service.CreateUser(name, email)
	if err != nil {
		fmt.Printf("could not create user %q - %q\n", name, email)
		return
	}
	fmt.Println("User created successfully")
	util.PrintUser(user)
}

func (interactiveCli *Interactive) UpdateUser() {
	id := util.ReadInput("Enter user ID: ")
	name := util.ReadInput("Enter new name: ")
	email := util.ReadInput("Enter new email: ")
	err := interactiveCli.service.UpdateUser(id, name, email)
	if err != nil {
		fmt.Printf("could not update user %q - %q - %q\n", id, name, email)
		return
	}
	fmt.Println("User updated successfully")
	user := domain.User{ID: id, Name: name, Email: email}
	util.PrintUser(user)
}

func (interactiveCli *Interactive) DeleteUser() {
	id := util.ReadInput("Enter user ID: ")
	err := interactiveCli.service.DeleteUser(id)
	if err != nil {
		fmt.Printf("could not delete user %q\n", id)
		return
	}
	fmt.Println("User deleted successfully")
}
