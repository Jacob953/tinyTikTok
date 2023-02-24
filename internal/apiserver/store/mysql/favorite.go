package mysql

import (
	"context"

	"gorm.io/gorm"
)

type favorites struct {
	db *gorm.DB
}

func (c favorites) Create(ctx context.Context, userId int64) error {
	return nil
}

func (c favorites) Delete(ctx context.Context, userId int64) error {
	return nil
}

func newFavorites(ds *datastore) *favorites {
	return &favorites{ds.db}
}
