package service

import (
	"context"

	"github.com/marmotedu/errors"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
)

// FeedSrv defines functions used to handle user request.
type FeedSrv interface {
	Get(ctx context.Context, latestTime int64) ([]*model.Video, int64, error)
}

var _ FeedSrv = (*feedService)(nil)

type feedService struct {
	store store.Factory
}

func newFeeds(s *service) FeedSrv {
	return &feedService{s.store}
}

func (f feedService) Get(ctx context.Context, latestTime int64) ([]*model.Video, int64, error) {
	videos, err := f.store.Videos().Feed(ctx, latestTime)
	if err != nil {
		return nil, -1, errors.WithCode(code.ErrDatabase, err.Error())
	}

	if len(videos) == 0 {
		return nil, -1, errors.WithCode(code.ErrDatabase, err.Error())
	}

	timestamp := videos[len(videos)-1].CreateAt

	return videos, timestamp, nil
}
