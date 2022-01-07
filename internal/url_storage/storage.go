package url_storage

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type UrlStorage struct {
	Addr       string
	Expiration time.Duration

	ctx    context.Context
	client *redis.Client
}

func (s *UrlStorage) Connect() error {
	s.ctx = context.TODO()
	s.client = redis.NewClient(&redis.Options{
		Addr: s.Addr,
	})

	err := s.client.Ping(s.ctx).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *UrlStorage) Get(hash string) (string, error) {
	val, err := s.client.Get(s.ctx, hash).Result()

	switch err {
	case nil:
		return val, nil
	case redis.Nil:
		return "", nil
	default:
		return "", err
	}
}

func (s *UrlStorage) Set(hash string, url string) error {
	err := s.client.Set(s.ctx, hash, url, s.Expiration).Err()
	if err != nil {
		return err
	}

	return nil
}
