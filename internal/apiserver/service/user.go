package service

import (
	"context"
	"regexp"

	"github.com/marmotedu/errors"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"
	genericapiserver "github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/server"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
)

// UserSrv defines functions used to handle user request.
type UserSrv interface {
	Get(ctx context.Context, userId int64) (*model.User, error)
	Register(ctx context.Context, username string, password string) (*model.User, error)
}

var _ UserSrv = (*userService)(nil)

type userService struct {
	store store.Factory
}

func newUsers(srv *service) *userService {
	return &userService{store: srv.store}
}

func (u userService) Get(ctx context.Context, userId int64) (*model.User, error) {
	user, err := u.store.Users().GetById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userService) Login(ctx context.Context, username string, password string) (*model.User, error) {
	user, err := u.store.Users().Get(ctx, username)
	if err != nil || password != user.Password {
		return nil, err
	}
	return user, nil
}

func (u userService) Register(ctx context.Context, username string, password string) (*model.User, error) {
	var err error

	_, err = u.store.Users().Get(ctx, username)
	if err == nil {
		match, _ := regexp.MatchString("Duplicate entry '.*' for key 'idx_name'", err.Error())
		if match {
			return nil, errors.WithCode(code.ErrUserAlreadyExist, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	user := &model.User{
		Id:       genericapiserver.SnowflakeSrv.NextID(),
		Name:     username,
		Password: password,
	}

	err = u.store.Users().Create(ctx, user)
	if err != nil {
		match, _ := regexp.MatchString("Duplicate entry '.*' for key 'idx_name'", err.Error())
		if match {
			return nil, errors.WithCode(code.ErrUserAlreadyExist, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}
	return user, nil
}
