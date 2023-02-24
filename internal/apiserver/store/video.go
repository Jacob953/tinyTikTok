package store

import (
	"context"

	model "github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"
)

type VideoStore interface {
	Create(ctx context.Context, video *model.Video) error
	Update(ctx context.Context, video *model.Video) error
	Get(ctx context.Context, videoId int64) (*model.Video, error)
	FavoriteList(ctx context.Context, userId int64) ([]*model.Video, error)
	PublishList(ctx context.Context, userId int64) ([]*model.Video, error)
	Feed(ctx context.Context, lastTime int64) ([]*model.Video, error)
}
