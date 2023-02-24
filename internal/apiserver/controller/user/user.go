package user

import (
	srv "github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/service"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
)

// NewUserController creates a user handler.
func NewUserController(store store.Factory) *UserController {
	return &UserController{
		srv: srv.NewService(store),
	}
}

// UserController create a user handler used to handle request for user resource.
type UserController struct {
	srv srv.Service
}

type UserRequest struct {
	Username string `protobuf:"bytes,1,req,name=username" form:"username,omitempty"` // 注册用户名，最长32个字符
	Password string `protobuf:"bytes,2,req,name=password" form:"password,omitempty"` // 密码，最长32个字符
}

type UserResponse struct {
	Id    int64  `protobuf:"varint,3,req,name=id"   json:"id,omitempty"`    // 用户id
	Token string `protobuf:"bytes,4,req,name=token" json:"token,omitempty"` // 用户鉴权token
}
