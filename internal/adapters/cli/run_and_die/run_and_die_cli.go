package run_and_die

import (
	"fmt"
	"gohexarc/cmd/registry"
	"gohexarc/internal/adapters/cli/util"
	"gohexarc/internal/domain"
	"io"
	"os"
)

type RunAndDie struct {
	services *registry.Services
	out      io.Writer
}

func NewRunAndDieCLI(services *registry.Services, out io.Writer) *RunAndDie {
	return &RunAndDie{services: services, out: out}
}

func (runAndDieCli *RunAndDie) Run() {
	if len(os.Args) == 2 {
		util.PrintUsage(runAndDieCli.out)
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
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func (runAndDieCli *RunAndDie) ListUsers() {
	users, err := runAndDieCli.services.UserService.ListUsers()
	if err != nil {
		err = fmt.Errorf("could not list users: %v", err)
		return
	}
	for _, user := range users {
		util.PrintUser(runAndDieCli.out, user)
	}
}

func (runAndDieCli *RunAndDie) GetUser() {
	user, err := runAndDieCli.services.UserService.GetUser(os.Args[3])
	if err != nil {
		errorMessage := fmt.Errorf("user %q not found", os.Args[3])
		fmt.Println(errorMessage)
		return
	}
	util.PrintUser(runAndDieCli.out, user)
}

func (runAndDieCli *RunAndDie) CreateUser() {
	user, err := runAndDieCli.services.UserService.CreateUser(os.Args[3], os.Args[4])
	if err != nil {
		errorMessage := fmt.Errorf("could not create user %q - %q", os.Args[3], os.Args[4])
		fmt.Println(errorMessage)
	}
	fmt.Println("User created successfully")
	util.PrintUser(runAndDieCli.out, user)
}

func (runAndDieCli *RunAndDie) UpdateUser() {
	err := runAndDieCli.services.UserService.UpdateUser(os.Args[3], os.Args[4], os.Args[5])
	if err != nil {
		errorMessage := fmt.Errorf("could not update user %q - %q - %q", os.Args[3], os.Args[4], os.Args[5])
		fmt.Println(errorMessage)
	}
	fmt.Println("User updated successfully")
	user := domain.User{ID: os.Args[3], Name: os.Args[4], Email: os.Args[5]}
	util.PrintUser(runAndDieCli.out, user)
}

func (runAndDieCli *RunAndDie) DeleteUser() {
	err := runAndDieCli.services.UserService.DeleteUser(os.Args[3])
	if err != nil {
		errorMessage := fmt.Errorf("could not delete user %q", os.Args[3])
		fmt.Println(errorMessage)
	}
	fmt.Println("User deleted successfully")
}
