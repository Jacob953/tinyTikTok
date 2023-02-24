package feed

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"

	"github.com/CSU-Apple-Lab/tinyTikTok/pkg/core"
)

type FeedRequest struct {
	LatestTime int64  `protobuf:"varint,1,opt,name=latest_time,json=latestTime" form:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string `protobuf:"bytes,2,opt,name=token"                        form:"token,omitempty"`       // 可选参数，登录用户设置
}

type FeedResponse struct {
	VideoList []*model.Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list,omitempty"` // 视频列表
	NextTime  int64          `protobuf:"varint,4,opt,name=next_time,json=nextTime"  json:"next_time,omitempty"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

func (f *FeedController) Get(ctx *gin.Context) {
	log.L(ctx).Info("feed get function called.")

	var r FeedRequest

	if err := ctx.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if errs := r.Validate(); len(errs) != 0 {
		core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	if r.LatestTime == 0 {
		r.LatestTime = time.Now().Unix()
	}

	vlst, timestamp, err := f.srv.Feeds().Get(ctx, r.LatestTime)
	if err != nil {
		core.WriteResponse(ctx, err, nil)

		return
	}

	core.WriteResponse(ctx, nil, FeedResponse{
		VideoList: vlst,
		NextTime:  timestamp,
	})
}
