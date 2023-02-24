package publish

import (
	srv "github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/service"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
)

func NewPublishController(store store.Factory) *PublishController {
	return &PublishController{
		srv: srv.NewService(store),
	}
}

// PublishController create a user handler used to handle request for user resource.
type PublishController struct {
	srv srv.Service
}
