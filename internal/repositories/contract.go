package repositories

import (
	"context"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
)

type Userer interface {
	GetAllUser(ctx context.Context) ([]entity.User, error)
}

type Todoer interface {
	GetAllTodo(ctx context.Context) ([]entity.Todo, error)
	CreateTodo(ctx context.Context, param interface{}) error
	GetOneTodo(ctx context.Context, param interface{}) (*entity.Todo, error)
	GetTodoStatus(ctx context.Context, param interface{}) (*entity.Todo, error)
	UpdateTodo(ctx context.Context, param interface{}, where interface{}) error
	DeleteTodo(ctx context.Context, param interface{}) error
}
