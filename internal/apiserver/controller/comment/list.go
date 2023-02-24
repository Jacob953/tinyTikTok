package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"
	"github.com/CSU-Apple-Lab/tinyTikTok/pkg/core"
)

type CommentListResponse struct {
	CommentList []*model.Comment `protobuf:"bytes,3,rep,name=comment_list,json=commentList" json:"comment_list,omitempty"` // 评论列表
}

func (c *CommentController) List(ctx *gin.Context) {
	log.L(ctx).Info("comment list function called.")

	var r CommentRequest

	if err := ctx.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if errs := r.Validate(); len(errs) != 0 {
		core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	clst, err := c.srv.Comments().List(ctx, r.VideoId)
	if err != nil {
		core.WriteResponse(ctx, err, nil)

		return
	}

	core.WriteResponse(ctx, nil, CommentListResponse{
		CommentList: clst,
	})
}
