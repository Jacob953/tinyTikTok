package service

import (
	"context"

	"github.com/marmotedu/errors"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"
)

type CommentSrv interface {
	Action(ctx context.Context, comment *model.Comment, actionType int32) (*model.Comment, error)
	List(ctx context.Context, videoId int64) ([]*model.Comment, error)
}

var _ CommentSrv = (*commentService)(nil)

type commentService struct {
	store store.Factory
}

func newComments(srv *service) *commentService {
	return &commentService{store: srv.store}
}

func (c commentService) Action(ctx context.Context, comment *model.Comment, actionType int32) (*model.Comment, error) {
	var com = &model.Comment{}
	var err error

	switch actionType {
	case 1:
		com, err = c.store.Comments().Create(ctx, comment)
		if err != nil {
			return nil, errors.WithCode(code.ErrDatabase, err.Error())
		}
	case 2:
		com, err = c.store.Comments().Delete(ctx, comment)
		if err != nil {
			return nil, errors.WithCode(code.ErrDatabase, err.Error())
		}
	}

	return com, nil
}

func (c commentService) List(ctx context.Context, videoId int64) ([]*model.Comment, error) {
	coms, err := c.store.Comments().List(ctx, videoId)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return coms, nil
}
