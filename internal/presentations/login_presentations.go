package presentations

import (
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"time"
)

type LoginPresentation struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginOutput struct {
	entity.User
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}
