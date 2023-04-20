package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/helper"
)

type todoGroup struct {
	db *sql.DB
}

func NewTodoGroup(db *sql.DB) TodoGrouper {
	return &todoGroup{db: db}
}

func (r *todoGroup) GetTodoGroupId(ctx context.Context, param interface{}) (*entity.TodoGroup, error) {
	result := &entity.TodoGroup{}

	wheres, vals := helper.QueryWhere(param)

	q := fmt.Sprintf(`
	SELECT
    	id
    FROM todo_groups %s ORDER BY created_at DESC  LIMIT 1`, wheres)

	err := r.db.QueryRowContext(ctx, q, vals...).Scan(
		&result.ID,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return result, err
}

func (r *todoGroup) GetAllTodoGroup(ctx context.Context, param interface{}) ([]entity.TodoGroup, error) {
	result := []entity.TodoGroup{}

	wheres, vals := helper.QueryWhere(param)

	q := fmt.Sprintf(`
	SELECT
    	id, title, unique_name,created_at, updated_at
    FROM todo_groups %s ORDER BY created_at DESC`, wheres)

	rows, err := r.db.QueryContext(ctx, q, vals...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var t entity.TodoGroup
		//err = rows.Scan(&usr)
		err = rows.Scan(&t.ID, &t.Title, &t.Unique, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, err
}

func (r *todoGroup) CreateTodoGroup(ctx context.Context, param interface{}) error {
	fields, vals := helper.QueryInsert(param)

	q := fmt.Sprintf(`INSERT INTO todo_groups %s`, fields)

	_, err := r.db.ExecContext(ctx, q, vals...)
	if err != nil {
		return err
	}

	return nil
}

//
//func (r *todoGroup) UpdateTodo(ctx context.Context, param interface{}, where interface{}) error {
//	fields, vals := helper.QueryUpdate(param, where)
//
//	q := fmt.Sprintf(`UPDATE todos SET %s`, fields)
//
//	_, err := r.db.ExecContext(ctx, q, vals...)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (r *todoGroup) DeleteTodo(ctx context.Context, param interface{}) error {
//	fields, vals := helper.QueryWhere(param)
//
//	q := fmt.Sprintf(`DELETE FROM todos %s`, fields)
//
//	_, err := r.db.ExecContext(ctx, q, vals...)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
