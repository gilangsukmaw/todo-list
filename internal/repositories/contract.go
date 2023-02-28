package repositories

import (
	"context"
	"go-fiber-v1/internal/entity"
)

type Userer interface {
	GetAllUser(ctx context.Context) ([]entity.User, error)
}
