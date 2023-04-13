package todo_group

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

type allTodoGroups struct {
	todoGroup repositories.TodoGrouper
}

func NewAllTodoGroup(
	todoGroup repositories.TodoGrouper,
) contract.UseCase {
	return &allTodoGroups{todoGroup: todoGroup}
}

func (u *allTodoGroups) Serve(dctx *fiber.Ctx, cfg *cfg.Config) server.Response {
	var (
		ctx    = dctx.Context()
		userId = fmt.Sprintf(`%v`, dctx.Locals("user_id"))
		//logger
		lf = logger.Field{
			EventName: "ucase get all todo group",
		}
		log = logger.NewLoggerField(lf)
	)

	logrus.WithField("event", log).Info("ucase_get_all_todo")

	data, err := u.todoGroup.GetTodoGroupId(ctx, entity.TodoGroup{UserId: userId})

	if err != nil {
		logrus.WithField("event",
			lf.Append("get todos got error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	logrus.WithField("event", log).Info("ucase get all todo success")

	return server.Response{Code: 200, Data: data, Message: "ok"}
}
