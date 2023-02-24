package store

//go:generate mockgen -self_package=github.com/Apple-Lab/tinyTikTok/internal/apiserver/store -destination mock_store.go -package store github.com/Apple-Lab/tinyTikTok/internal/apiserver/store Factory,UserStore,VideoStore,CommentStore

// Factory defines the iam platform storage interface.
type Factory interface {
	Users() UserStore
	Videos() VideoStore
	Comments() CommentStore
	Favorites() FavoriteStore
	Close() error
}

var client Factory

// Client return the store client instance.
func Client() Factory {
	return client
}

// SetClient set the iam store client.
func SetClient(factory Factory) {
	client = factory
}
