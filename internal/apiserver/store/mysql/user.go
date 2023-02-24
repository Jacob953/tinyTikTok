package mysql

import (
	"context"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"

	"github.com/marmotedu/errors"
	"gorm.io/gorm"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"
)

type users struct {
	db *gorm.DB
}

func (u *users) Get(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}

	err := u.db.Where("name = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUserNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return user, nil
}

func (u *users) GetById(ctx context.Context, userId int64) (*model.User, error) {
	user := &model.User{}
	err := u.db.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUserNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}
	return user, nil
}

func (u *users) Create(ctx context.Context, user *model.User) error {
	return u.db.Create(user).Error
}

func (u *users) Update(ctx context.Context, user *model.User) error {
	return u.db.Save(user).Error
}

func newUsers(ds *datastore) *users {
	return &users{ds.db}
}
