package todo_group

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	cfg "gitlab.com/todo-list-app1/todo-list-backend/cfg/env"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/helper"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/presentations"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/repositories"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/contract"
	"gitlab.com/todo-list-app1/todo-list-backend/lib/logger"
)

type createTodoGroup struct {
	todoGroupRepo repositories.TodoGrouper
}

func NewCreateTodoGroup(
	todoGroupRepo repositories.TodoGrouper,
) contract.UseCase {
	return &createTodoGroup{
		todoGroupRepo: todoGroupRepo,
	}
}

func (u *createTodoGroup) Serve(dctx *fiber.Ctx, cfg *cfg.Config) server.Response {
	var (
		ctx    = dctx.Context()
		param  = presentations.CreateTodoGroupParam{}
		userId = fmt.Sprintf(`%v`, dctx.Locals("user_id"))
		err    = dctx.BodyParser(&param)
		//logger
		lf = logger.Field{
			EventName: "ucase create todo group",
		}
		log = logger.NewLoggerField(lf)
	)

	logrus.WithField("event", log).Info("ucase create todo")

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

	err = u.todoGroupRepo.CreateTodoGroup(ctx, entity.TodoGroup{
		ID:     id.String(),
		UserId: userId,
		Title:  param.Title,
		Unique: helper.GenerateUniqueName(),
	})

	if err != nil {
		logrus.WithField("event",
			lf.Append("save todo error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	return server.Response{Code: 201, Message: "Created"}
}
