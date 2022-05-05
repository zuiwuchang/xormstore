package redis

import (
	"github.com/go-redis/redis/v8"
)

// 定義默認參數值
var defaultOptions = options{
	client:   nil,
	redisURL: `redis://@localhost:6789/?dial_timeout=3&db=0&read_timeout=6s&max_retries=2`,
}

type options struct {
	client   *redis.Client
	redisURL string
	options  *redis.Options
}

type Option interface {
	apply(*options)
}

type funcOption struct {
	f func(*options)
}

func (fdo *funcOption) apply(do *options) {
	fdo.f(do)
}
func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}
func WithClient(client *redis.Client) Option {
	return newFuncOption(func(o *options) {
		o.client = client
	})
}
func WithOptions(opts *redis.Options) Option {
	return newFuncOption(func(o *options) {
		o.options = opts
	})
}
func WithURL(redisURL string) Option {
	return newFuncOption(func(o *options) {
		o.redisURL = redisURL
	})
}
