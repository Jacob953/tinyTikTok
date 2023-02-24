package publish

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/server"
	"github.com/CSU-Apple-Lab/tinyTikTok/pkg/core"
)

type PublishRequest struct {
	Token string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"` // 用户鉴权token
	Data  []byte `protobuf:"bytes,2,req,name=data"  json:"data,omitempty"`  // 视频数据
	Title string `protobuf:"bytes,3,req,name=title" json:"title,omitempty"` // 视频标题
}

func (c *PublishController) Action(ctx *gin.Context) {
	log.L(ctx).Info("comment action function called.")

	var r PublishRequest

	err := ctx.ShouldBindQuery(&r)
	if err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if errs := r.Validate(); len(errs) != 0 {
		core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	err = c.srv.Publishs().Action(ctx, &model.Video{
		Id:    server.SnowflakeSrv.NextID(),
		Title: r.Title,
	}, r.Data)
	if err != nil {
		core.WriteResponse(ctx, err, nil)

		return
	}

	core.WriteResponse(ctx, nil, nil)
}
