package service

import (
	"context"
	"os"
	"os/exec"
	"strconv"

	"github.com/marmotedu/errors"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"
)

type PublishSrv interface {
	Action(ctx context.Context, user *model.Video, data []byte) error
	List(ctx context.Context, userId int64) ([]*model.Video, error)
}

var _ PublishSrv = (*publishService)(nil)

type publishService struct {
	store store.Factory
}

func newPublishs(srv *service) *publishService {
	return &publishService{store: srv.store}
}

func (p publishService) Action(ctx context.Context, video *model.Video, data []byte) error {
	var err error
	fileName := strconv.FormatInt(video.Id, 10)
	videoPath := fileName + ".mp4"
	coverPath := fileName + ".jpg"

	err = os.WriteFile("temp.mp4", data, 0644)
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	cmd := exec.Command(
		"ffmpeg",
		"-i",
		"temp.mp4",
		"-c:v",
		"libx264",
		"-c:a",
		"aac",
		"-b:a",
		"192k",
		"-vf",
		"scale=640:-2",
		"-preset",
		"slow",
		videoPath,
	)
	err = cmd.Run()
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	cmd = exec.Command("ffmpeg", "-i", videoPath, "-ss", "00:00:00", "-vframes", "1", coverPath)
	err = cmd.Run()
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	video.PlayUrl = videoPath
	video.CoverUrl = coverPath
	err = p.store.Videos().Create(ctx, video)
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (p publishService) List(ctx context.Context, userId int64) ([]*model.Video, error) {
	plst, err := p.store.Videos().PublishList(ctx, userId)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return plst, nil
}
