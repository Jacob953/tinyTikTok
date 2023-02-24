package service

import (
	"context"

	"github.com/marmotedu/errors"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"
)

type FavoriteSrv interface {
	Action(ctx context.Context, videoId int64, actionType int32) error
	List(ctx context.Context, userId int64) ([]*model.Video, error)
}

var _ FavoriteSrv = (*favoriteService)(nil)

type favoriteService struct {
	store store.Factory
}

func newFavorites(srv *service) *favoriteService {
	return &favoriteService{store: srv.store}
}

func (f favoriteService) List(ctx context.Context, userId int64) ([]*model.Video, error) {
	flst, err := f.store.Videos().FavoriteList(ctx, userId)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return flst, nil
}

func (f favoriteService) Action(ctx context.Context, videoId int64, actionType int32) error {
	var err error

	video, err := f.store.Videos().Get(ctx, videoId)
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	switch actionType {
	case 1:
		video.FavoriteCount += 1
		err = f.store.Videos().Update(ctx, video)
		if err != nil {
			return errors.WithCode(code.ErrDatabase, err.Error())
		}
	case 2:
		video.FavoriteCount -= 1
		err = f.store.Videos().Update(ctx, video)
		if err != nil {
			return errors.WithCode(code.ErrDatabase, err.Error())
		}
	}

	return nil
}
