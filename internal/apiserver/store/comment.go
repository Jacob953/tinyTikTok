package store

import (
	"context"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"
)

type CommentStore interface {
	Create(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	Delete(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	List(ctx context.Context, videoId int64) ([]*model.Comment, error)
}
