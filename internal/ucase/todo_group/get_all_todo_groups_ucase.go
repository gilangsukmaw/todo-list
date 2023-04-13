package todo_group

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	cfg "gitlab.com/todo-list-app1/todo-list-backend/cfg/env"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/presentations"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/repositories"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/contract"
	"gitlab.com/todo-list-app1/todo-list-backend/lib/logger"
)

type allTodoGroups struct {
	todoRepo  repositories.Todoer
	todoGroup repositories.TodoGrouper
}

func NewAllTodoGroup(
	todoRepo repositories.Todoer,
	todoGroup repositories.TodoGrouper,
) contract.UseCase {
	return &allTodoGroups{todoRepo: todoRepo, todoGroup: todoGroup}
}

func (u *allTodoGroups) Serve(dctx *fiber.Ctx, cfg *cfg.Config) server.Response {
	var (
		ctx    = dctx.Context()
		userId = fmt.Sprintf(`%v`, dctx.Locals("user_id"))
		resp   = []presentations.AllTodoGroupResponse{}
		//logger
		lf = logger.Field{
			EventName: "ucase get all todo group",
		}
		log = logger.NewLoggerField(lf)
	)

	logrus.WithField("event", log).Info("ucase_get_all_todo")

	data, err := u.todoGroup.GetAllTodoGroup(ctx, entity.TodoGroup{UserId: userId})

	if err != nil {
		logrus.WithField("event",
			lf.Append("get todos got error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	for _, todoGroup := range data {
		tmp := presentations.AllTodoGroupResponse{}

		todos, err := u.todoRepo.CountTodo(ctx, entity.Todo{GroupId: todoGroup.ID})

		if err != nil {
			logrus.WithField("event",
				lf.Append("count todos got error", fmt.Sprintf(`%s`, err))).Error()
			return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
		}

		tmp.ID = todoGroup.ID
		tmp.Title = todoGroup.Title
		tmp.UniqueName = todoGroup.Unique
		tmp.TodoTotal = todos

		resp = append(resp, tmp)
	}

	logrus.WithField("event", log).Info("ucase get all todo success")

	return server.Response{Code: 200, Data: resp, Message: "ok"}
}
