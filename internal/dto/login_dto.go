package dto

import (
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/presentations"
)

func LoginDTO(user *entity.User, token string) presentations.LoginOutput {
	var (
		result = presentations.LoginOutput{}
	)

	result.ID = user.ID
	result.Email = user.Email
	result.Username = user.Username
	result.Role = user.Role
	result.Avatar = user.Avatar
	result.Token = token

	return result
}
