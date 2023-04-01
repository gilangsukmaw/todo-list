package todos

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	cfg "gitlab.com/todo-list-app1/todo-list-backend/cfg/env"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/consts"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/repositories"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/contract"
	"gitlab.com/todo-list-app1/todo-list-backend/lib/logger"
)

type undoneTodo struct {
	todoRepo repositories.Todoer
}

func NewUndoneTodos(
	todoRepo repositories.Todoer,
) contract.UseCase {
	return &undoneTodo{todoRepo: todoRepo}
}

func (u *undoneTodo) Serve(dctx *fiber.Ctx, cfg *cfg.Config) server.Response {
	var (
		ctx = dctx.Context()
		//logger
		lf = logger.Field{
			EventName: "ucase todo undone",
		}
		id  = dctx.Params("id")
		log = logger.NewLoggerField(lf)
	)

	logrus.WithField("event", log).Info("ucase todo undone")

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

	if todo.Status == consts.TodoOnProgress {
		logrus.WithField("event",
			lf.Append("todo status is already on progress", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 422, Message: "todo status already on-progress"}
	}

	if err != nil {
		logrus.WithField("event",
			lf.Append("save todo error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	result, err := u.todoRepo.GetOneTodo(ctx, entity.Todo{ID: id})

	if err != nil {
		logrus.WithField("event",
			lf.Append("get todos got error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	logrus.WithField("event", log).Info("ucase todo undone")

	return server.Response{Code: 200, Data: result, Message: "ok"}
}
