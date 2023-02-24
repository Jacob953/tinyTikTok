package service

import (
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
)

// Service defines functions used to return resource interface.
type Service interface {
	Feeds() FeedSrv
	Users() UserSrv
	Publishs() PublishSrv
	Favorites() FavoriteSrv
	Comments() CommentSrv
}

type service struct {
	store store.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Feeds() FeedSrv {
	return newFeeds(s)
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}

func (s *service) Publishs() PublishSrv {
	return newPublishs(s)
}

func (s *service) Favorites() FavoriteSrv {
	return newFavorites(s)
}

func (s *service) Comments() CommentSrv {
	return newComments(s)
}
