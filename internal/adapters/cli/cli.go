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

func RunCli(service port.UserService) {
	var cli CLI

	if len(os.Args) == 2 {
		cli = interactiveCli.NewInteractiveCLI(service)
	} else {
		cli = run_and_die.NewRunAndDieCLI(service)
	}

	cli.Run()
}
