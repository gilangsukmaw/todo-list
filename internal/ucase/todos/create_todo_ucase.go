package todos

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/todo-list-app1/todo-list-backend/cfg/yaml"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/helper"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/presentations"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/repositories"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/contract"
	"gitlab.com/todo-list-app1/todo-list-backend/lib/logger"
)

type createTodo struct {
	todoRepo repositories.Todoer
}

func NewCreateTodo(
	todoRepo repositories.Todoer,
) contract.UseCase {
	return &createTodo{todoRepo: todoRepo}
}

func (u *createTodo) Serve(dctx *fiber.Ctx, cfg *yaml.Config) server.Response {
	var (
		ctx   = dctx.Context()
		param = presentations.CreateTodoParam{}
		err   = dctx.BodyParser(&param)
		//logger
		lf = logger.Field{
			EventName: "ucase create todo",
		}
		log = logger.NewLoggerField(lf)
	)

	logrus.WithField("event", log).Info("ucase create todo")

	//time.Sleep(3 * time.Second)

	if err != nil {
		logrus.WithField("event",
			lf.Append("error casting request body", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: "error casting request body"}
	}

	if param.Title == "" {
		logrus.WithField("event",
			lf.Append("request body validation error", fmt.Sprintf(`title is null or empty`))).Error()
		return server.Response{Code: 400, Message: "title cannot be empty"}
	}

	id := uuid.New()

	err = u.todoRepo.CreateTodo(ctx, entity.Todo{
		ID:     id.String(),
		UserId: id.String(),
		Title:  param.Title,
		Color:  helper.RandomColor(),
		//Status: "done",
	})

	if err != nil {
		logrus.WithField("event",
			lf.Append("save todo error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	todo, err := u.todoRepo.GetOneTodo(ctx, entity.Todo{ID: id.String()})

	if err != nil {
		logrus.WithField("event",
			lf.Append("get todo got error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	return server.Response{Code: 201, Data: todo, Message: "Created"}
}
