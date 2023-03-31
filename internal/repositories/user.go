package repositories

import (
	"context"
	"database/sql"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
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
