package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Damisicode/social/internal/store"
	"github.com/go-redis/redis/v8"
)

type PostStore struct {
	rdb *redis.Client
}

const PostExpTime = time.Minute

func (p *PostStore) Get(ctx context.Context, postID int64) (*store.Post, error) {
	cacheKey := fmt.Sprintf("post-%v", postID)

	data, err := p.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var post store.Post
	if data != "" {
		err := json.Unmarshal([]byte(data), &post)
		if err != nil {
			return nil, err
		}
	}

	return &post, nil
}

func (p *PostStore) Set(ctx context.Context, post *store.Post) error {
	if post.ID == 0 {
		return fmt.Errorf("missing post ID")
	}
	cacheKey := fmt.Sprintf("post-%v", post.ID)

	json, err := json.Marshal(post)
	if err != nil {
		return err
	}

	return p.rdb.SetEX(ctx, cacheKey, json, PostExpTime).Err()
}
