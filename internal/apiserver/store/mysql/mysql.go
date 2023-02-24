package mysql

import (
	"fmt"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/logger"
	genericoptions "github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/options"
	"github.com/CSU-Apple-Lab/tinyTikTok/pkg/db"

	"sync"

	"github.com/marmotedu/errors"
	"gorm.io/gorm"
)

var (
	mysqlFactory store.Factory
	once         sync.Once
)

// GetMySQLFactory create mysql factory with the given configs.
func GetMySQLFactory(opts *genericoptions.MySQLOptions) (store.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory")
	}
	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		options := &db.Options{
			Host:                  opts.Host,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			LogLevel:              opts.LogLevel,
			Logger:                logger.New(opts.LogLevel),
		}
		dbIns, err = db.New(options)

		mysqlFactory = &datastore{dbIns}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}

func (ds *datastore) Videos() store.VideoStore {
	return newVideos(ds)
}

func (ds *datastore) Comments() store.CommentStore {
	return newComments(ds)
}

func (ds *datastore) Favorites() store.FavoriteStore {
	return newFavorites(ds)
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}

	return db.Close()
}
