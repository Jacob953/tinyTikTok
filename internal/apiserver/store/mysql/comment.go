package mysql

import (
	"context"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"

	"github.com/marmotedu/errors"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"

	"gorm.io/gorm"
)

type comments struct {
	db *gorm.DB
}

func (c comments) Create(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	var err error

	tx := c.db.Begin()
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	comment.FormatDate()

	err = c.db.Create(comment).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	video, err := store.Client().Videos().Get(ctx, comment.VideoId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	video.CommentCount += 1
	err = store.Client().Videos().Update(ctx, video)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return comment, nil
}

func (c comments) Delete(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	var err error

	tx := c.db.Begin()
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	comment.FormatDate()

	err = c.db.Where("comment_id = ?", comment.Id).Delete(&model.Comment{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	video, err := store.Client().Videos().Get(ctx, comment.VideoId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	video.CommentCount -= 1
	err = store.Client().Videos().Update(ctx, video)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return comment, nil
}

func (c comments) List(ctx context.Context, videoId int64) ([]*model.Comment, error) {
	var coms []*model.Comment

	d := c.db.Where(" video_id = ?", videoId).
		Order("create_date desc").
		Find(&coms)

	return coms, d.Error
}

func newComments(ds *datastore) *comments {
	return &comments{ds.db}
}
