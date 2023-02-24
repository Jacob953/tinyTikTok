package mysql

import (
	"context"
	"time"

	"github.com/marmotedu/errors"
	"gorm.io/gorm"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"

	model "github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"
)

type videos struct {
	db *gorm.DB
}

func (v videos) Feed(ctx context.Context, lastTime int64) ([]*model.Video, error) {
	var ret []*model.Video

	d := v.db.Where("create_at < ?", time.Unix(lastTime, 0)).
		Order("create_at desc").
		Limit(30).
		Preload("Author").
		Find(&ret)

	return ret, d.Error
}

func (v videos) PublishList(ctx context.Context, userId int64) ([]*model.Video, error) {
	var plst []*model.Video

	d := v.db.Table("video").
		Where("user_id = ?", userId).
		Find(&plst)

	return plst, d.Error
}

func (v videos) Create(ctx context.Context, video *model.Video) error {
	return v.db.Create(&video).Error
}

func (v videos) Update(ctx context.Context, video *model.Video) error {
	return v.db.Save(video).Error
}

func (v videos) Get(ctx context.Context, videoId int64) (*model.Video, error) {
	video := &model.Video{}

	err := v.db.Where("video_id = ?", videoId).First(&video).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUserNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return video, nil
}

func (v videos) FavoriteList(ctx context.Context, userId int64) ([]*model.Video, error) {
	var flst []*model.Video

	d := v.db.Table("video").
		Joins("INNER JOIN likes ON favorite.video_id = video.video_id").
		Where("favorite.user_id = ?", userId).
		Find(&flst)

	return flst, d.Error
}

func newVideos(ds *datastore) *videos {
	return &videos{ds.db}
}
