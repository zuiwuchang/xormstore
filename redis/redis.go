package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/zuiwuchang/xormstore/encode"
)

type Store struct {
	client *redis.Client
}

func New(opt ...Option) (store *Store, e error) {
	opts := defaultOptions
	for _, o := range opt {
		o.apply(&opts)
	}
	var client *redis.Client
	if opts.client != nil {
		client = opts.client
	} else {
		o := opts.options
		if o == nil {
			o, e = redis.ParseURL(opts.redisURL)
			if e != nil {
				return
			}
		}
		client = redis.NewClient(o)
	}
	store = &Store{
		client: client,
	}
	return
}
func (s *Store) Close() error {
	return s.client.Close()
}
func (s *Store) Del(key string) error {
	cmd := s.client.Del(context.Background(), key)
	return cmd.Err()
}
func (s *Store) Get(key string) (interface{}, error) {
	cmd := s.client.Get(context.Background(), key)
	val, e := cmd.Result()
	if e != nil {
		return nil, e
	}
	var to interface{}
	e = encode.GobDecode([]byte(val), &to)
	if e != nil {
		return nil, e
	}
	return to, nil
}
func (s *Store) Put(key string, value interface{}) error {
	val, e := encode.GobEncode(value)
	if e != nil {
		return e
	}
	cmd := s.client.Set(context.Background(), key, val, 0)
	return cmd.Err()
}
