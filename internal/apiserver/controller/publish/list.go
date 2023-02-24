package publish

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"
	"github.com/CSU-Apple-Lab/tinyTikTok/pkg/core"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"
)

type PublishListRequest struct {
	UserId int64  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"` // 用户id
	Token  string `protobuf:"bytes,2,req,name=token"                json:"token,omitempty"`   // 用户鉴权token
}

type PublishListResponse struct {
	VideoList []*model.Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list,omitempty"` // 用户发布的视频列表
}

func (c *PublishController) List(ctx *gin.Context) {
	log.L(ctx).Info("comment list function called.")

	var r PublishListRequest

	if err := ctx.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if errs := r.Validate(); len(errs) != 0 {
		core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	plst, err := c.srv.Publishs().List(ctx, r.UserId)
	if err != nil {
		core.WriteResponse(ctx, err, nil)

		return
	}

	core.WriteResponse(ctx, nil, PublishListResponse{
		VideoList: plst,
	})
}
