package entity

import (
	"time"
)

type TodoGroup struct {
	ID        string     `db:"id" json:"id,omitempty"`
	UserId    string     `db:"user_id" json:"user_id,omitempty"`
	Title     string     `db:"title" json:"title,omitempty"`
	Color     *string    `db:"color" json:"color,omitempty"`
	Unique    string     `db:"unique_name" json:"unique_name,omitempty"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

type Todo struct {
	ID        string     `db:"id" json:"id,omitempty"`
	GroupId   string     `db:"group_id" json:"group_id,omitempty"`
	UserId    string     `db:"user_id" json:"user_id"`
	Title     string     `db:"title" json:"title,omitempty"`
	Status    string     `db:"status" json:"status,omitempty"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt,omitempty"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt,omitempty"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
