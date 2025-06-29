package interactiveCli

import (
	"bufio"
	"fmt"
	"gohexarc/internal/adapters/cli/util"
	"gohexarc/internal/domain"
	"gohexarc/internal/port"
	"io"
)

type Interactive struct {
	service port.UserService
	in      *bufio.Reader
	out     io.Writer
}

func NewInteractiveCLI(service port.UserService, in io.Reader, out io.Writer) *Interactive {
	return &Interactive{service: service, in: bufio.NewReader(in), out: out}
}

func (interactiveCli *Interactive) Run() {
	for {
		util.PrintUsage(interactiveCli.out)
		command := util.ReadInput("Choose one of the available options above: ", interactiveCli.in, interactiveCli.out)
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
		util.PrintUser(interactiveCli.out, user)
	}
}

func (interactiveCli *Interactive) GetUser() {
	id := util.ReadInput("Enter user ID: ", interactiveCli.in, interactiveCli.out)
	user, err := interactiveCli.service.GetUser(id)
	if err != nil {
		fmt.Printf("user %q not found\n", id)
		return
	}
	util.PrintUser(interactiveCli.out, user)
}

func (interactiveCli *Interactive) CreateUser() {
	name := util.ReadInput("Enter name: ", interactiveCli.in, interactiveCli.out)
	email := util.ReadInput("Enter email: ", interactiveCli.in, interactiveCli.out)
	user, err := interactiveCli.service.CreateUser(name, email)
	if err != nil {
		fmt.Printf("could not create user %q - %q\n", name, email)
		return
	}
	fmt.Println("User created successfully")
	util.PrintUser(interactiveCli.out, user)
}

func (interactiveCli *Interactive) UpdateUser() {
	id := util.ReadInput("Enter user ID: ", interactiveCli.in, interactiveCli.out)
	name := util.ReadInput("Enter new name: ", interactiveCli.in, interactiveCli.out)
	email := util.ReadInput("Enter new email: ", interactiveCli.in, interactiveCli.out)
	err := interactiveCli.service.UpdateUser(id, name, email)
	if err != nil {
		fmt.Printf("could not update user %q - %q - %q\n", id, name, email)
		return
	}
	fmt.Println("User updated successfully")
	user := domain.User{ID: id, Name: name, Email: email}
	util.PrintUser(interactiveCli.out, user)
}

func (interactiveCli *Interactive) DeleteUser() {
	id := util.ReadInput("Enter user ID: ", interactiveCli.in, interactiveCli.out)
	err := interactiveCli.service.DeleteUser(id)
	if err != nil {
		fmt.Printf("could not delete user %q\n", id)
		return
	}
	fmt.Fprintln(interactiveCli.out, "User deleted successfully")
}
