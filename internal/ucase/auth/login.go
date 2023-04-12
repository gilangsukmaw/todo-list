package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	cfg "gitlab.com/todo-list-app1/todo-list-backend/cfg/env"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/dto"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/helper"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/presentations"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/repositories"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/contract"
	"gitlab.com/todo-list-app1/todo-list-backend/lib/logger"
)

type login struct {
	userRepo repositories.Userer
}

func NewLogin(
	userRepo repositories.Userer,
) contract.UseCase {
	return &login{
		userRepo: userRepo,
	}
}

func (u *login) Serve(dctx *fiber.Ctx, cfg *cfg.Config) server.Response {
	var (
		param = presentations.LoginPresentation{}
		err   = dctx.BodyParser(&param)
		//logger
		lf = logger.Field{
			EventName: "ucase_login",
		}
		log = logger.NewLoggerField(lf)
	)

	logrus.WithField("event", log).Info("ucase_login")

	if err != nil {
		logrus.WithField("event",
			lf.Append("cast body got error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: "error cast request body"}
	}

	//find user
	user, err := u.userRepo.GetOneUser(dctx.Context(), entity.User{Email: param.Email})
	if err != nil {
		logrus.WithField("event",
			lf.Append("get user got", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: err}
	}

	if user == nil {
		logrus.WithField("event",
			lf.Append("get user got", fmt.Sprintf(`nil`))).Error()
		return server.Response{Code: 401, Message: "username or password is invalid"}
	}

	ok := helper.CheckPasswordHash(user.Password, param.Password)
	if !ok {
		logrus.WithField("event",
			lf.Append("password not valid", fmt.Sprintf(``))).Error()
		return server.Response{Code: 401, Message: "username or password is invalid"}
	}

	tokenString, expiredAt, err := helper.GenerateJWT(user, cfg.JwtToken, cfg.JwtExpired)
	if err != nil {
		logrus.WithField("event",
			lf.Append("generate token got", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: "error cast request body"}
	}

	logrus.WithField("event", log).Info("succes login")

	return server.Response{Code: 200, Data: dto.LoginDTO(user, tokenString, expiredAt), Message: "ok"}
}
