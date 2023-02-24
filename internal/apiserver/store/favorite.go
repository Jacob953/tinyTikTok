package store

import (
	"context"
)

type FavoriteStore interface {
	Create(ctx context.Context, userId int64) error
	Delete(ctx context.Context, userId int64) error
}
