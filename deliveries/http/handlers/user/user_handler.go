package handlers

import service "github.com/dezh-tech/panda/services/user"

type User struct {
	service service.User
}

func NewUserService(userSvc service.User) User {
	return User{
		service: userSvc,
	}
}
