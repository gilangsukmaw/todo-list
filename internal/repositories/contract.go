package repositories

import (
	"context"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
)

type Userer interface {
	GetAllUser(ctx context.Context) ([]entity.User, error)
	GetOneUser(ctx context.Context, param interface{}) (*entity.User, error)
}

type Todoer interface {
	GetAllTodo(ctx context.Context, param interface{}) ([]entity.Todo, error)
	CreateTodo(ctx context.Context, param interface{}) error
	GetOneTodo(ctx context.Context, param interface{}) (*entity.Todo, error)
	GetTodoStatus(ctx context.Context, param interface{}) (*entity.Todo, error)
	UpdateTodo(ctx context.Context, param interface{}, where interface{}) error
	DeleteTodo(ctx context.Context, param interface{}) error
	CountTodo(ctx context.Context, param interface{}) (int, error)
}

type TodoGrouper interface {
	GetTodoGroupId(ctx context.Context, param interface{}) (*entity.TodoGroup, error)
	GetAllTodoGroup(ctx context.Context, param interface{}) ([]entity.TodoGroup, error)
	CreateTodoGroup(ctx context.Context, param interface{}) error
}
