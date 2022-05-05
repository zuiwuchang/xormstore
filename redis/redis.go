package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Put(key string, value interface{}) error
// Get(key string) (interface{}, error)
// Del(key string) error

// func NewLevelDBStore(dbfile string) (*LevelDBStore, error) {
// 	db := &LevelDBStore{}
// 	h, err := leveldb.OpenFile(dbfile, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	db.store = h
// 	return db, nil
// }
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
	cmd.Result()
	return cmd.Err()
}
