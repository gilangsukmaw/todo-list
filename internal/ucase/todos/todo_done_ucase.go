package todos

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gitlab.com/todo-list-app1/todo-list-backend/cfg/yaml"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/consts"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/repositories"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/ucase/contract"
	"gitlab.com/todo-list-app1/todo-list-backend/lib/logger"
)

type doneTodo struct {
	todoRepo repositories.Todoer
}

func NewDoneTodos(
	todoRepo repositories.Todoer,
) contract.UseCase {
	return &doneTodo{todoRepo: todoRepo}
}

func (u *doneTodo) Serve(dctx *fiber.Ctx, cfg *yaml.Config) server.Response {
	var (
		ctx = dctx.Context()
		//logger
		lf = logger.Field{
			EventName: "ucase todo done",
		}
		id  = dctx.Params("id")
		log = logger.NewLoggerField(lf)
	)

	logrus.WithField("event", log).Info("ucase todo done")

	todo, err := u.todoRepo.GetTodoStatus(ctx, entity.Todo{ID: id})

	if err != nil {
		logrus.WithField("event",
			lf.Append("get todos got error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: fmt.Sprintf(`%s`, err)}
	}

	fmt.Println("todo status nih gan --> ", todo)

	if todo == nil {
		logrus.WithField("event",
			lf.Append("get todos got nil", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 404, Message: "not found, please check the id you've sent!"}
	}

	if todo.Status == consts.TodoDone {
		logrus.WithField("event",
			lf.Append("todo status already done", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 422, Message: "todo status already done"}
	}

	err = u.todoRepo.UpdateTodo(ctx, entity.Todo{Status: "done"}, entity.Todo{ID: id})

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

	logrus.WithField("event", log).Info("ucase todo done")

	return server.Response{Code: 200, Data: result, Message: "ok"}
}
