package cli

import (
	"gohexarc/cmd/registry"
	"gohexarc/internal/adapters/cli/interactive"
	"gohexarc/internal/adapters/cli/run_and_die"
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

type FactoryCLI func(services *registry.Services) CLI

func DefaultCLIFactory(services *registry.Services) CLI {
	if len(os.Args) == 2 {
		return interactive.NewInteractiveCLI(services.UserService, os.Stdin, os.Stdout)
	}
	return run_and_die.NewRunAndDieCLI(services.UserService, os.Stdout)
}

func RunCli(services *registry.Services, factory FactoryCLI) {
	cli := factory(services)
	cli.Run()
}
