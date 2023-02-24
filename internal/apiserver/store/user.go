package store

import (
	"context"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"
)

type UserStore interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Get(ctx context.Context, username string) (*model.User, error)
	GetById(ctx context.Context, userId int64) (*model.User, error)
}
