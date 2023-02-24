package favorite

import (
	srv "github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/service"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
)

func NewFavoriteController(store store.Factory) *FavoriteController {
	return &FavoriteController{
		srv: srv.NewService(store),
	}
}

// FavoriteController create a user handler used to handle request for user resource.
type FavoriteController struct {
	srv srv.Service
}
