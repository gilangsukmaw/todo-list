package entity

import (
	"time"
)

type Todo struct {
	ID        string     `db:"id" json:"id,omitempty"`
	UserId    string     `db:"user_id" json:"user_id"`
	Title     string     `db:"title" json:"title,omitempty"`
	Status    string     `db:"status" json:"status,omitempty"`
	Color     string     `db:"color" json:"color,omitempty"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt,omitempty"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt,omitempty"`
	DeletedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`
}
