package cache

import (
	"context"

	"github.com/Damisicode/social/internal/store"
	"github.com/go-redis/redis/v8"
)

type Storage struct {
	Users interface {
		Get(context.Context, int64) (*store.User, error)
		Set(context.Context, *store.User) error
	}
	Posts interface {
		Get(context.Context, int64) (*store.Post, error)
		Set(context.Context, *store.Post) error
	}
}

func NewRedisStorage(rdb *redis.Client) Storage {
	return Storage{
		Users: &UserStore{rdb: rdb},
		Posts: &PostStore{rdb: rdb},
	}
}
