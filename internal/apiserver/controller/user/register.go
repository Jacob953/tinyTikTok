package user

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/auth"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/code"
	"github.com/CSU-Apple-Lab/tinyTikTok/pkg/core"
)

// Register add new user to the storage.
func (u *UserController) Register(ctx *gin.Context) {
	log.L(ctx).Info("get user function called.")

	var r UserRequest

	err := ctx.ShouldBindQuery(&r)
	if err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	r.Password, _ = auth.Encrypt(r.Password)
	user, err := u.srv.Users().Register(ctx, r.Username, r.Password)
	if err != nil {
		core.WriteResponse(ctx, err, nil)

		return
	}

	core.WriteResponse(ctx, nil, UserResponse{
		Id:    user.Id,
		Token: "",
	})
}
