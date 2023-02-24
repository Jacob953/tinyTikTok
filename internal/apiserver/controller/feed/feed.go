package feed

import (
	srv "github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/service"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
)

func NewFeedController(store store.Factory) *FeedController {
	return &FeedController{
		srv: srv.NewService(store),
	}
}

// FeedController create a user handler used to handle request for user resource.
type FeedController struct {
	srv srv.Service
}
