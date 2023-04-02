package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/helper"
)

type user struct {
	db *sql.DB
}

func NewUser(db *sql.DB) Userer {
	return &user{db: db}
}

func (r *user) GetAllUser(ctx context.Context) ([]entity.User, error) {
	result := []entity.User{}

	rows, err := r.db.QueryContext(ctx, "SELECT id, username, full_name, email, avatar, role FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var usr entity.User
		//err = rows.Scan(&usr)
		err = rows.Scan(&usr.ID, &usr.Username, &usr.Fullname, &usr.Email, &usr.Avatar, &usr.Role)
		if err != nil {
			return nil, err
		}
		result = append(result, usr)
	}

	return result, err
}

func (r *user) GetOneUser(ctx context.Context, param interface{}) (*entity.User, error) {
	usr := entity.User{}

	wheres, vals := helper.QueryWhere(param)

	q := fmt.Sprintf(`SELECT id, username, full_name, email, avatar, role, password FROM users %s LIMIT 1`, wheres)

	fmt.Println("q --> ", q)
	fmt.Println("wheres --> ", wheres)
	fmt.Println("val --> ", vals)

	err := r.db.QueryRowContext(ctx, q, vals...).Scan(&usr.ID, &usr.Username, &usr.Fullname, &usr.Email, &usr.Avatar, &usr.Role, &usr.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &usr, nil
}
