package registry

import (
	"gohexarc/internal/adapters/repository/memory"
	"gohexarc/internal/adapters/repository/sqlite"
	"gohexarc/internal/port"
	"gohexarc/internal/service"
	"gohexarc/internal/util"
)

type Services struct {
	UserService port.UserService
}

func NewServices() *Services {
	var userRepository port.UserRepository
	if util.ShouldUseInMemory() {
		userRepository = memory.NewInMemoryUserRepository()
	} else {
		db, _ := util.OpenDatabase()
		userRepository = sqlite.NewSqliteUserRepository(db)
	}
	userService := service.NewUserService(userRepository)

	return &Services{UserService: userService}
}
