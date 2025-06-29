package cli

import (
	interactiveCli "gohexarc/internal/adapters/cli/interactive"
	"gohexarc/internal/adapters/cli/run_and_die"
	"gohexarc/internal/port"
	"os"
)

type CLI interface {
	Run()
	ListUsers()
	GetUser()
	CreateUser()
	UpdateUser()
	DeleteUser()
}

type FactoryCLI func(service port.UserService) CLI

func DefaultCLIFactory(service port.UserService) CLI {
	if len(os.Args) == 2 {
		return interactiveCli.NewInteractiveCLI(service, os.Stdin, os.Stdout)
	}
	return run_and_die.NewRunAndDieCLI(service, os.Stdout)
}

func RunCli(service port.UserService, factory FactoryCLI) {
	cli := factory(service)
	cli.Run()
}
