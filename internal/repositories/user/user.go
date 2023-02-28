package repositories

import (
	"context"
	"database/sql"
	"go-fiber-v1/internal/entity"
	"go-fiber-v1/internal/repositories"
)

type user struct {
	db *sql.DB
}

func NewUser(db *sql.DB) repositories.Userer {
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
