package user

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-v1/cfg/yaml"
	"go-fiber-v1/internal/repositories"
	"go-fiber-v1/internal/server"
	"go-fiber-v1/internal/ucase/contract"
)

type users struct {
	userRepo repositories.Userer
}

func NewUsers(
	userRepo repositories.Userer,
) contract.UseCase {
	return &users{
		userRepo: userRepo,
	}
}

func (u *users) Serve(ctx *fiber.Ctx, cfg *yaml.Config) server.Response {

	userList, err := u.userRepo.GetAllUser(ctx.Context())
	if err != nil {
		return server.Response{Code: 500, Message: err}
	}

	return server.Response{Code: 200, Data: userList, Message: "ok"}
}
