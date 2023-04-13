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

type allTodos struct {
	todoRepo      repositories.Todoer
	todoGroupRepo repositories.TodoGrouper
}

func NewAllTodos(
	todoRepo repositories.Todoer,
	todoGroupRepo repositories.TodoGrouper,
) contract.UseCase {
	return &allTodos{
		todoRepo:      todoRepo,
		todoGroupRepo: todoGroupRepo,
	}
}

func (u *allTodos) Serve(dctx *fiber.Ctx, cfg *cfg.Config) server.Response {
	var (
		ctx    = dctx.Context()
		userId = fmt.Sprintf(`%v`, dctx.Locals("user_id"))
		unique = dctx.Params("unique")
		//logger
		lf = logger.Field{
			EventName: "ucase get all todo",
		}
		log = logger.NewLoggerField(lf)
	)

	logrus.WithField("event", log).Info("ucase_get_all_todo")

	if unique == "" {
		logrus.WithField("event",
			lf.Append("get todos got error", "unique is empty"))
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, "unique is empty")}
	}

	todoGroup, err := u.todoGroupRepo.GetTodoGroupId(ctx, entity.TodoGroup{UserId: userId, Unique: unique})

	if err != nil {
		logrus.WithField("event",
			lf.Append("get todos got error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	todos, err := u.todoRepo.GetAllTodo(ctx, entity.Todo{UserId: userId, GroupId: todoGroup.ID})

	if err != nil {
		logrus.WithField("event",
			lf.Append("get todos got error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	logrus.WithField("event", log).Info("ucase get all todo success")

	return server.Response{Code: 200, Data: todos, Message: "ok"}
}
