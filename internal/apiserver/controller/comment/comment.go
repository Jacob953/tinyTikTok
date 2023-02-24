package comment

import (
	srv "github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/service"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
)

func NewCommentController(store store.Factory) *CommentController {
	return &CommentController{
		srv: srv.NewService(store),
	}
}

// CommentController create a user handler used to handle request for user resource.
type CommentController struct {
	srv srv.Service
}

type CommentRequest struct {
	Token   string `protobuf:"bytes,1,req,name=token"                  form:"token,omitempty"`    // 用户鉴权token
	VideoId int64  `protobuf:"varint,2,req,name=video_id,json=videoId" form:"video_id,omitempty"` // 视频id
}
