package todos

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gitlab.com/todo-list-app1/todo-list-backend/cfg/yaml"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/repositories"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/contract"
	"gitlab.com/todo-list-app1/todo-list-backend/lib/logger"
)

type allTodos struct {
	todoRepo repositories.Todoer
}

func NewAllTodos(
	todoRepo repositories.Todoer,
) contract.UseCase {
	return &allTodos{todoRepo: todoRepo}
}

func (u *allTodos) Serve(dctx *fiber.Ctx, cfg *yaml.Config) server.Response {
	var (
		ctx = dctx.Context()
		//logger
		lf = logger.Field{
			EventName: "ucase get all todo",
		}
		log = logger.NewLoggerField(lf)
	)

	logrus.WithField("event", log).Info("ucase_get_all_todo")

	todos, err := u.todoRepo.GetAllTodo(ctx)

	if err != nil {
		logrus.WithField("event",
			lf.Append("get todos got error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	logrus.WithField("event", log).Info("ucase get all todo success")

	return server.Response{Code: 200, Data: todos, Message: "ok"}
}
