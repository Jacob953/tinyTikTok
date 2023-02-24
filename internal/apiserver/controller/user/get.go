package user

import (
	"github.com/gin-gonic/gin"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"

	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"

	"github.com/CSU-Apple-Lab/tinyTikTok/pkg/core"
)

type UserGetRequest struct {
	Id    int64  `protobuf:"varint,1,req,name=id"   json:"id,omitempty"`    // 用户id
	Token string `protobuf:"bytes,2,req,name=token" json:"token,omitempty"` // 用户鉴权token
}

type UserGetResponse struct {
	User *model.User `protobuf:"bytes,3,req,name=user" json:"user,omitempty"` // 用户信息
}

// Get add new user to the storage.
func (u *UserController) Get(ctx *gin.Context) {
	log.L(ctx).Info("get user function called.")

	var r UserGetRequest

	err := ctx.ShouldBindQuery(&r)
	if err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	user, err := u.srv.Users().Get(ctx, r.Id)
	if err != nil {
		core.WriteResponse(ctx, err, nil)

		return
	}

	core.WriteResponse(ctx, nil, UserGetResponse{
		User: user,
	})
}
