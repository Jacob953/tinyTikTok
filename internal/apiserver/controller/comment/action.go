package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"

	"github.com/CSU-Apple-Lab/tinyTikTok/pkg/core"
)

type CommentCreateRequest struct {
	CommentRequest
	ActionType  int32  `protobuf:"varint,3,req,name=action_type,json=actionType"  form:"action_type,omitempty"`  // 1-发布评论，2-删除评论
	CommentText string `protobuf:"bytes,4,opt,name=comment_text,json=commentText" form:"comment_text,omitempty"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   int64  `protobuf:"varint,5,opt,name=comment_id,json=commentId"    form:"comment_id,omitempty"`   // 要删除的评论id，在action_type=2的时候使用
}

type CommentCreateResponse struct {
	Comment *model.Comment `protobuf:"bytes,3,opt,name=comment" json:"comment,omitempty"` // 评论成功返回评论内容，不需要重新拉取整个列表
}

func (c *CommentController) Action(ctx *gin.Context) {
	log.L(ctx).Info("comment action function called.")

	var r CommentCreateRequest

	err := ctx.ShouldBindQuery(&r)
	if err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if errs := r.Validate(); len(errs) != 0 {
		core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	cm, err := c.srv.Comments().Action(ctx, &model.Comment{
		Id:      r.CommentId,
		Content: r.CommentText,
		VideoId: r.VideoId,
	}, r.ActionType)

	if err != nil {
		core.WriteResponse(ctx, err, nil)

		return
	}

	core.WriteResponse(ctx, nil, CommentCreateResponse{
		Comment: cm,
	})
}
