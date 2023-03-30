package todos

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gitlab.com/todo-list-app1/todo-list-backend/cfg/yaml"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/repositories"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/contract"
	"gitlab.com/todo-list-app1/todo-list-backend/lib/logger"
)

type deleteTodo struct {
	todoRepo repositories.Todoer
}

func NewDeleteTodos(
	todoRepo repositories.Todoer,
) contract.UseCase {
	return &deleteTodo{todoRepo: todoRepo}
}

func (u *deleteTodo) Serve(dctx *fiber.Ctx, cfg *yaml.Config) server.Response {
	var (
		ctx = dctx.Context()
		//logger
		lf = logger.Field{
			EventName: "ucase todo delete",
		}
		id  = dctx.Params("id")
		log = logger.NewLoggerField(lf)
	)

	logrus.WithField("event", log).Info("ucase todo delete")

	todo, err := u.todoRepo.GetTodoStatus(ctx, entity.Todo{ID: id})

	err = u.todoRepo.UpdateTodo(ctx, entity.Todo{Status: "on-progress"}, entity.Todo{ID: id})

	if err != nil {
		logrus.WithField("event",
			lf.Append("get todos got error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	if todo == nil {
		logrus.WithField("event",
			lf.Append("get todos got nil", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 404, Message: "not found, please check the id you've sent!"}
	}

	err = u.todoRepo.DeleteTodo(ctx, entity.Todo{ID: id})

	if err != nil {
		logrus.WithField("event",
			lf.Append("delete todos got error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	logrus.WithField("event", log).Info("ucase todo delete")

	return server.Response{Code: 200, Data: true, Message: "ok"}
}
