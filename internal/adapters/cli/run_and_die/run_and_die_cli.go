package run_and_die

import (
	"fmt"
	"gohexarc/internal/adapters/cli/util"
	"gohexarc/internal/domain"
	"gohexarc/internal/port"
	"os"
)

type Interactive struct {
	service port.UserService
}

func NewRunAndDieCLI(service port.UserService) *Interactive {
	return &Interactive{service: service}
}

func (runAndDieCli *Interactive) Run() {
	if len(os.Args) == 2 {
		util.PrintUsage()
		return
	}

	command := os.Args[2]
	switch command {
	case util.List:
		runAndDieCli.ListUsers()
	case util.Get:
		runAndDieCli.GetUser()
	case util.Create:
		runAndDieCli.CreateUser()
	case util.Update:
		runAndDieCli.UpdateUser()
	case util.Delete:
		runAndDieCli.DeleteUser()
	}
}

func (runAndDieCli *Interactive) ListUsers() {
	users, err := runAndDieCli.service.ListUsers()
	if err != nil {
		err = fmt.Errorf("could not list users: %v", err)
		return
	}
	for _, user := range users {
		util.PrintUser(user)
	}
}

func (runAndDieCli *Interactive) GetUser() {
	user, err := runAndDieCli.service.GetUser(os.Args[3])
	if err != nil {
		errorMessage := fmt.Errorf("user %q not found", os.Args[3])
		fmt.Println(errorMessage)
		return
	}
	util.PrintUser(user)
}

func (runAndDieCli *Interactive) CreateUser() {
	user, err := runAndDieCli.service.CreateUser(os.Args[3], os.Args[4])
	if err != nil {
		errorMessage := fmt.Errorf("could not create user %q - %q", os.Args[3], os.Args[4])
		fmt.Println(errorMessage)
	}
	fmt.Println("User created successfully")
	util.PrintUser(user)
}

func (runAndDieCli *Interactive) UpdateUser() {
	err := runAndDieCli.service.UpdateUser(os.Args[3], os.Args[4], os.Args[5])
	if err != nil {
		errorMessage := fmt.Errorf("could not update user %q - %q - %q", os.Args[3], os.Args[4], os.Args[5])
		fmt.Println(errorMessage)
	}
	fmt.Println("User updated successfully")
	user := domain.User{ID: os.Args[3], Name: os.Args[4], Email: os.Args[5]}
	util.PrintUser(user)
}

func (runAndDieCli *Interactive) DeleteUser() {
	err := runAndDieCli.service.DeleteUser(os.Args[3])
	if err != nil {
		errorMessage := fmt.Errorf("could not delete user %q", os.Args[3])
		fmt.Println(errorMessage)
	}
	fmt.Println("User deleted successfully")
}
