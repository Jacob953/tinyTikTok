package favorite

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"
	"github.com/CSU-Apple-Lab/tinyTikTok/pkg/core"
)

type FavoriteActionRequest struct {
	Token      string `protobuf:"bytes,1,req,name=token"                        json:"token,omitempty"`       // 用户鉴权token
	VideoId    int64  `protobuf:"varint,2,req,name=video_id,json=videoId"       json:"video_id,omitempty"`    // 视频id
	ActionType int32  `protobuf:"varint,3,req,name=action_type,json=actionType" json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}

func (c *FavoriteController) Action(ctx *gin.Context) {
	log.L(ctx).Info("comment action function called.")

	var r FavoriteActionRequest

	err := ctx.ShouldBindQuery(&r)
	if err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if errs := r.Validate(); len(errs) != 0 {
		core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	err = c.srv.Favorites().Action(ctx, r.VideoId, r.ActionType)
	if err != nil {
		core.WriteResponse(ctx, err, nil)

		return
	}

	core.WriteResponse(ctx, nil, nil)
}
