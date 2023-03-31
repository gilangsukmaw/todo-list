package todos

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	cfg "gitlab.com/todo-list-app1/todo-list-backend/cfg/env"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/repositories"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/contract"
	"gitlab.com/todo-list-app1/todo-list-backend/lib/logger"
)

type oneTodos struct {
	todoRepo repositories.Todoer
}

func NewOneTodos(
	todoRepo repositories.Todoer,
) contract.UseCase {
	return &oneTodos{todoRepo: todoRepo}
}

func (u *oneTodos) Serve(dctx *fiber.Ctx, cfg *cfg.Config) server.Response {
	var (
		ctx = dctx.Context()
		//logger
		lf = logger.Field{
			EventName: "ucase get one todo",
		}
		log = logger.NewLoggerField(lf)
	)

	logrus.WithField("event", log).Info("ucase_get_all_todo")

	todo, err := u.todoRepo.GetOneTodo(ctx, entity.Todo{ID: ""})

	if err != nil {
		logrus.WithField("event",
			lf.Append("get todos got error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	logrus.WithField("event", log).Info("ucase get one todo success")

	return server.Response{Code: 200, Data: todo, Message: "ok"}
}
