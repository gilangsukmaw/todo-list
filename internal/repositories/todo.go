package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/helper"
)

type todo struct {
	db *sql.DB
}

func NewTodo(db *sql.DB) Todoer {
	return &todo{db: db}
}

func (r *todo) GetAllTodo(ctx context.Context, param interface{}) ([]entity.Todo, error) {
	result := []entity.Todo{}

	wheres, vals := helper.QueryWhere(param)

	q := fmt.Sprintf(`
	SELECT
    	id,
    	group_id,
    	user_id,
    	title,
    	status,
    	created_at,
    	updated_at
    FROM todos %s ORDER BY created_at DESC`, wheres)

	rows, err := r.db.QueryContext(ctx, q, vals...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var t entity.Todo
		//err = rows.Scan(&usr)
		err = rows.Scan(&t.ID, &t.GroupId, &t.UserId, &t.Title, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, err
}

func (r *todo) GetOneTodo(ctx context.Context, param interface{}) (*entity.Todo, error) {
	result := &entity.Todo{}

	wheres, vals := helper.QueryWhere(param)

	q := fmt.Sprintf(`
	SELECT
    	id,
    	user_id,
    	title,
    	status,
    	created_at
    FROM todos %s ORDER BY created_at DESC  LIMIT 1`, wheres)

	err := r.db.QueryRowContext(ctx, q, vals...).Scan(
		&result.ID,
		&result.UserId,
		&result.Title,
		&result.Status,
		&result.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return result, err
}

func (r *todo) GetTodoStatus(ctx context.Context, param interface{}) (*entity.Todo, error) {
	result := &entity.Todo{}

	wheres, vals := helper.QueryWhere(param)

	q := fmt.Sprintf(`SELECT status FROM todos %s LIMIT 1`, wheres)

	err := r.db.QueryRowContext(ctx, q, vals...).Scan(&result.Status)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return result, err
}

func (r *todo) CreateTodo(ctx context.Context, param interface{}) error {
	fields, vals := helper.QueryInsert(param)

	q := fmt.Sprintf(`INSERT INTO todos %s`, fields)

	fmt.Println("ini q --> ", q)

	_, err := r.db.ExecContext(ctx, q, vals...)
	if err != nil {
		return err
	}

	return nil
}

func (r *todo) UpdateTodo(ctx context.Context, param interface{}, where interface{}) error {
	fields, vals := helper.QueryUpdate(param, where)

	q := fmt.Sprintf(`UPDATE todos SET %s`, fields)

	_, err := r.db.ExecContext(ctx, q, vals...)
	if err != nil {
		return err
	}

	return nil
}

func (r *todo) DeleteTodo(ctx context.Context, param interface{}) error {
	fields, vals := helper.QueryWhere(param)

	q := fmt.Sprintf(`DELETE FROM todos %s`, fields)

	_, err := r.db.ExecContext(ctx, q, vals...)
	if err != nil {
		return err
	}

	return nil
}

func (r *todo) CountTodo(ctx context.Context, param interface{}) (int, error) {
	var count int

	wheres, vals := helper.QueryWhere(param)

	q := fmt.Sprintf(`
	SELECT 
   		COUNT(id) 
    FROM todos %s`, wheres)

	rows, err := r.db.QueryContext(ctx, q, vals...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	if err != nil {
		return 0, err
	}

	return count, err
}
